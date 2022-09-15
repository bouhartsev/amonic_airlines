package repository

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

//language=MySQL
const getUserByIdQuery = `select
                     	    roleid, email, firstname, lastname, officeid, birthdate, active
                          from users
                          where id = ?`

func (r *repository) GetUserById(ctx context.Context, userId int) (*domain.User, error) {
	row := r.conn.QueryRowContext(ctx, getUserByIdQuery, userId)

	var user domain.User

	err := row.Scan(
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

	user.Id = &userId

	return &user, nil
}
