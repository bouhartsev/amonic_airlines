package repository

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
)

func (r *repository) GetUsers(ctx context.Context, request *domain.GetUsersRequest) ([]*domain.User, error) {
	//language=MySQL
	query := `select
                id, roleid, email, firstname, lastname, officeid, birthdate, active
                from users `

	args := make([]interface{}, 0)

	if request.OfficeId != nil {
		query += "where officeid = ?"
		args = append(args, *request.OfficeId)
	}

	rows, err := r.conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User

		err := rows.Scan(
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

		users = append(users, &user)
	}

	return users, nil
}
