package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

func (c *Core) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.User, error) {
	var userId int
	err := c.db.QueryRowContext(ctx, `select id from users where lower(email) = lower(?)`, request.Email).Scan(&userId)
	if err == nil {
		return nil, errdomain.EmailAlreadyTakenError
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.logger.Error(err.Error())
		return nil, errdomain.NewInternalError(err.Error())
	}

	result, err := c.db.ExecContext(
		ctx,
		`insert into users(roleid, email, firstname, lastname, officeid, birthdate, password, active)
                           values(2, ?, ?, ?, ?, cast(? as date), ?, true)`,
		request.Email,
		request.FirstName,
		request.LastName,
		request.OfficeId,
		request.Birthdate,
		request.Password,
	)

	if err != nil {
		return nil, err
	}

	insertedId, err := result.LastInsertId()

	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	user, err := c.GetUser(ctx, &domain.GetUserRequest{UserId: int(insertedId)})
	if err != nil {
		return nil, err
	}

	return user.User, nil
}

func (c *Core) GetUsers(ctx context.Context, request *domain.GetUsersRequest) (*domain.GetUsersResponse, error) {
	q := `select id, roleid, email, firstname, lastname, officeid, DATE_FORMAT(birthdate, "%Y-%m-%d"), timestampdiff(year, birthdate, now()), active from users `
	var args []any

	if request.OfficeId != nil {
		q += fmt.Sprintf("where officeid = ?")
		args = append(args, *request.OfficeId)
	}

	rows, err := c.db.QueryContext(ctx, q, args...)

	if err != nil {
		c.logger.Error(err.Error())
		return nil, errdomain.NewInternalError(err.Error())
	}

	var users []*domain.User

	for rows.Next() {
		var user domain.User

		err = rows.Scan(
			&user.Id,
			&user.RoleId,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.OfficeId,
			&user.Birthdate,
			&user.Age,
			&user.Active,
		)

		if err != nil {
			c.logger.Error(err.Error())
			return nil, errdomain.NewInternalError(err.Error())
		}

		users = append(users, &user)
	}

	return &domain.GetUsersResponse{Users: users}, nil
}

func (c *Core) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.GetUserResponse, error) {
	q := `select roleid, email, firstname, lastname, officeid, DATE_FORMAT(birthdate, "%Y-%m-%d"), timestampdiff(year, birthdate, now()), active from users where id = ?`

	row := c.db.QueryRowContext(ctx, q, request.UserId)

	var user domain.User

	err := row.Scan(
		&user.RoleId,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.OfficeId,
		&user.Birthdate,
		&user.Age,
		&user.Active,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errdomain.UserNotFoundError
		}
		c.logger.Error(err.Error())
		return nil, errdomain.NewInternalError(err.Error())
	}

	user.Id = &request.UserId

	return &domain.GetUserResponse{User: &user}, nil
}

func (c *Core) UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) error {
	var (
		args  []any
		query []string
	)

	if request.OfficeId != nil {
		query = append(query, "officeId = ?")
		args = append(args, *request.OfficeId)
	}

	if request.Email != nil {
		query = append(query, "email = ?")
		args = append(args, *request.Email)
	}

	if request.FirstName != nil {
		query = append(query, "firstName = ?")
		args = append(args, *request.FirstName)
	}

	if request.LastName != nil {
		query = append(query, "lastName = ?")
		args = append(args, *request.LastName)
	}

	if request.RoleId != nil {
		query = append(query, "roleId = ?")
		args = append(args, *request.RoleId)
	}

	arguments := strings.Join(query, ",")
	args = append(args, request.UserId)
	q := fmt.Sprintf("update users set %s where id = ?", arguments)

	_, err := c.db.ExecContext(ctx, q, args...)

	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}

	return nil
}

func (c *Core) GetUserLogins(ctx context.Context, id int) (*domain.GetUserLoginsResponse, error) {
	rows, err := c.db.QueryContext(ctx, `select DATE_FORMAT(loginTime, "%Y-%m-%d %H:%i"), DATE_FORMAT(logoutTime, "%Y-%m-%d %H:%i"), TIMEDIFF(logoutTime, loginTime), errorReason
                                               from user_logins where userId = ? and loginTime > date_sub(now(), interval 30 day)
                                               order by loginTime desc
                                               limit 200 offset 1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logins []domain.UserLogin

	for rows.Next() {
		var login domain.UserLogin
		if err = rows.Scan(
			&login.LoginTime,
			&login.LogoutTime,
			&login.TimeSpent,
			&login.Error,
		); err != nil {
			return nil, err
		}

		logins = append(logins, login)
	}

	resp := &domain.GetUserLoginsResponse{UserLogins: logins}

	if len(logins) > 0 {
		var crashes int
		err = c.db.QueryRowContext(ctx, "select count(*) from (select id from`user_logins` where userId = ? and logoutTime is null and loginTime > date_sub(now(), interval 30 day) limit 150 offset 1) t", id).Scan(&crashes)
		if err != nil {
			return nil, err
		}
		resp.NumberOfCrashes = crashes

		err = c.db.QueryRowContext(ctx, "select DATE_FORMAT(loginTime, \"%Y-%m-%d %H:%i\") from user_logins where userId = ? and logoutTime is null and confirmed = false and errorReason is null and loginTime > date_sub(now(), interval 30 day) order by loginTime desc limit 1 offset 1", id).Scan(&resp.LastLoginErrorDatetime)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	return resp, nil
}

func (c *Core) SwitchUserStatus(ctx context.Context, id int) error {
	_, err := c.db.ExecContext(ctx, `update users set active = !active where id = ?`, id)
	return err
}
