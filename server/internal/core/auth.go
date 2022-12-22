package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func (c *Core) SignIn(ctx context.Context, request *domain.SignInRequest) (*domain.SignInResponse, error) {
	var user domain.User

	row := c.db.QueryRowContext(ctx, `select id, IncorrectLoginTries, NextLoginTime from users where email = ?`, request.Login)

	err := row.Scan(
		&user.Id,
		&user.IncorrectLoginTries,
		&user.NextLoginTime,
	)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errdomain.InvalidCredentialsError
	}

	if user.NextLoginTime != nil && user.NextLoginTime.After(time.Now()) &&
		user.IncorrectLoginTries != nil && *user.IncorrectLoginTries >= 3 {
		return nil, &errdomain.ErrorResponse{
			Message: "Series of incorrect credentials.",
			Type:    errdomain.InvalidRequestType,
			Code:    "invalid_credentials:series",
			Details: struct {
				NextTry time.Time
			}{
				NextTry: *user.NextLoginTime,
			},
		}
	}

	//language=MySQL
	const signInQuery = `select
                     		id, roleid, email, firstname, lastname, officeid, birthdate, active
                     	 from users
                     	 where email = ?
                     	 and password = ?`

	row = c.db.QueryRowContext(ctx, signInQuery, request.Login, request.Password)

	err = row.Scan(
		&user.Id,
		&user.RoleId,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.OfficeId,
		&user.Birthdate,
		&user.Active,
	)

	// incorrect credentials
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		_, err = c.db.ExecContext(ctx, `update users set IncorrectLoginTries = IncorrectLoginTries + 1, NextLoginTime = DATE_ADD(NOW(), INTERVAL 10 SECOND) where email = ?`, request.Login)

		if err != nil {
			c.logger.Error(err.Error())
			return nil, err
		}

		return nil, errdomain.InvalidCredentialsError
	}

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	if user.Active != nil && *user.Active == false {
		return nil, errdomain.UserDisabledError
	}

	claims := domain.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // 30 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: &user,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(c.cfg.TokenKey))

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	_, err = c.db.ExecContext(ctx, `update users set NextLoginTime = NULL, IncorrectLoginTries = 0 where id = ?`, user.Id)
	if err != nil {
		return nil, err
	}

	_, err = c.db.ExecContext(ctx, "insert into `user_logins`(userId) values(?)", user.Id)
	if err != nil {
		return nil, err
	}

	return &domain.SignInResponse{Token: token, User: user}, nil
}

func (c *Core) SignOut(ctx context.Context) error {
	token, err := c.GetTokenFromContext(ctx)
	if err != nil {
		return err
	}

	var lastUserLoginId int
	err = c.db.QueryRowContext(ctx, "select id from `user_logins` where userId = ? and logoutTime IS NULL and errorReason IS NULL", token.User.Id).Scan(&lastUserLoginId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errdomain.NoActiveLoginsDetectedError
		}
		return err
	}

	_, err = c.db.ExecContext(ctx, "update `user_logins` set logoutTime = NOW(), confirmed = true where id = ?", lastUserLoginId)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) ReportLastLogoutError(ctx context.Context, req *domain.ReportLastLogoutErrorRequest) error {
	token, err := c.GetTokenFromContext(ctx)
	if err != nil {
		return err
	}

	var message string

	if req.Reason != "" {
		message = req.Reason
	} else if req.SoftwareCrash {
		message = "Software crash"
	} else {
		message = "System crash"
	}

	var lastLogoutId int
	err = c.db.QueryRowContext(ctx, "select id from `user_logins` where userId = ? and logoutTime IS NULL and confirmed = false order by loginTime desc limit 1 offset 1", token.User.Id).Scan(&lastLogoutId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("nothing to report")
		}
		return err
	}

	_, err = c.db.ExecContext(ctx, "update `user_logins` set errorReason = ?, confirmed = true where id = ?", message, lastLogoutId)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) GetTokenFromContext(ctx context.Context) (*domain.AuthClaims, error) {
	ctxToken := ctx.Value("token")

	if ctxToken == nil {
		return nil, errdomain.ErrTokenIsNotPresented
	}

	accessToken, ok := ctxToken.(string)
	if !ok {
		return nil, errdomain.ErrInvalidAccessToken
	}

	accessToken = strings.Split(accessToken, "Bearer ")[1]

	token, err := jwt.ParseWithClaims(accessToken, &domain.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(c.cfg.TokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.AuthClaims); ok && token.Valid {
		claims.RawToken = accessToken
		return claims, nil
	}

	return nil, errdomain.ErrInvalidAccessToken
}

func (c *Core) SetExpiredTokenError(ctx context.Context, userId int) error {
	var loginId int
	err := c.db.QueryRowContext(ctx, "select id from user_logins where userId = ? and logoutTime is null and errorReason is null order by loginTime desc limit 1", userId).Scan(&loginId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	_, err = c.db.ExecContext(ctx, "update user_logins set errorReason = ?, confirmed = true where id = ?", "Expired auth token", loginId)
	if err != nil {
		return err
	}

	return nil
}
