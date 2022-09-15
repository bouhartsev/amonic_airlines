package repository

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

//language=MySQL
const signInQuery = `select
                     	id, roleid, email, firstname, lastname, officeid, birthdate, active
                     from users
                     where email = ?
                     and password = ?`

func (r *repository) SignIn(ctx context.Context, request *domain.SignInRequest) (*domain.User, error) {
	row := r.conn.QueryRowContext(ctx, signInQuery, request.Login, request.Password)

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
		return nil, err
	}

	return &user, nil
}
