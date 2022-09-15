package repository

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

//language=MySQL
const getUserByEmailQuery = `select
                     	       id, roleid, firstname, lastname, officeid, birthdate, active
                             from users
                             where email = ?`

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.conn.QueryRowContext(ctx, getUserByEmailQuery, email)

	var user domain.User

	err := row.Scan(
		&user.Id,
		&user.RoleId,
		&user.FirstName,
		&user.LastName,
		&user.OfficeId,
		&user.Birthdate,
		&user.Active,
	)

	if err != nil {
		return nil, err
	}

	user.Email = &email

	return &user, nil
}
