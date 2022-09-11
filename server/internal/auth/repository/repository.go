package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bouhartsev/amonic_airlines/server/internal/auth"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain/errdomain"
)

type repository struct {
	conn *sql.DB
}

func NewAuthRepository(conn *sql.DB) auth.Repository {
	return &repository{conn: conn}
}

func (r *repository) SignIn(ctx context.Context, request *domain.SignInRequest) (*domain.User, error) {
	//language=MySQL
	const query = `
select
	id, roleid, email, firstname, lastname, officeid, birthdate, active
from users
where email = ?
and password = ?
`

	row := r.conn.QueryRowContext(ctx, query, request.Login, request.Password)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var user domain.User

	err := row.Scan(
		&user.Id,
		&user.RoleId,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.OfficeId,
		&user.Birthdate,
		&user.Active,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errdomain.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}
