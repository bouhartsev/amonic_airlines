package repository

import (
	"context"
	"github.com/bouhartsev/amonic_airlines/server/internal/domain"
	"github.com/bouhartsev/amonic_airlines/server/pkg/ptr"
)

//language=MySQL
const createUserQuery = `insert into users(roleid, email, firstname, lastname, officeid, birthdate, password, active)
                         values(2, ?, ?, ?, ?, cast(? as date), ?, false)`

func (r *repository) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.User, error) {

	result, err := r.conn.ExecContext(ctx, createUserQuery,
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
		return nil, err
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
