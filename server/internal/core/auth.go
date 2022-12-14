package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)), // 30 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: &user,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(c.cfg.TokenKey))

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	go func(db *sql.DB, id int) {
		_, err = c.db.ExecContext(ctx, `update users set NextLoginTime = NULL, IncorrectLoginTries = 0 where id = ?`, id)

		if err != nil {
			c.logger.Error(err.Error())
		}
	}(c.db, *user.Id)

	return &domain.SignInResponse{Token: token}, nil
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
		return claims, nil
	}

	return nil, errdomain.ErrInvalidAccessToken
}