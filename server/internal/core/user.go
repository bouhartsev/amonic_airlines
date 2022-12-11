package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
	"github.com/bouhartsev/amonic_airlines/server/pkg/ptr"
)

func (c *Core) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.User, error) {
	_, err := c.db.QueryContext(ctx, `select id from users where email = ?`, request.Email)

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
                           values(2, ?, ?, ?, ?, cast(? as date), ?, false)`,
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

	return &domain.User{
		Id:        ptr.Int(int(insertedId)),
		RoleId:    ptr.Int(2),
		Email:     ptr.String(request.Email),
		FirstName: ptr.String(request.FirstName),
		LastName:  ptr.String(request.LastName),
		OfficeId:  ptr.Int(request.OfficeId),
		Birthdate: ptr.Time(request.Birthdate),
		Active:    ptr.Bool(false),
	}, nil
}

func (c *Core) GetUsers(ctx context.Context, request *domain.GetUsersRequest) (*domain.GetUsersResponse, error) {
	q := `select id, roleid, email, firstname, lastname, officeid, timestampdiff(year, birthdate, now()), active from users `
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
	q := `select roleid, email, firstname, lastname, officeid, timestampdiff(year, birthdate, now()), active from users where id = ?`

	row := c.db.QueryRowContext(ctx, q, request.UserId)

	var user domain.User

	err := row.Scan(
		&user.RoleId,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.OfficeId,
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
