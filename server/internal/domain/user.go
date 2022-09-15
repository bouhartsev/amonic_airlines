package domain

import "time"

type User struct {
	Id        *int       `json:"id,omitempty"`
	RoleId    *int       `json:"roleId,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Password  *string    `json:"-"`
	FirstName *string    `json:"firstName,omitempty"`
	LastName  *string    `json:"lastName,omitempty"`
	OfficeId  *int       `json:"officeId,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Active    *bool      `json:"active,omitempty"`
}

type CreateUserRequest struct {
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	OfficeId  int       `json:"officeId,omitempty"`
	Birthdate time.Time `json:"birthdate,omitempty"`
	Password  string    `json:"password,omitempty"`
}

type GetUserRequest struct {
	UserId int `json:"userId"`
}

type GetUsersRequest struct {
	OfficeId *int `json:"officeId,omitempty"`
}

type UpdateUserRequest struct {
	RoleId *int `json:"roleId,omitempty"`
}
