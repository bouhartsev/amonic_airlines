package domain

import "time"

type User struct {
	Id                  *int       `json:"id,omitempty"`
	RoleId              *int       `json:"roleId,omitempty"`
	Email               *string    `json:"email,omitempty"`
	Password            *string    `json:"-"`
	FirstName           *string    `json:"firstName,omitempty"`
	LastName            *string    `json:"lastName,omitempty"`
	OfficeId            *int       `json:"officeId,omitempty"`
	Birthdate           *time.Time `json:"birthdate,omitempty"`
	Age                 *int       `json:"age,omitempty"`
	Active              *bool      `json:"active,omitempty"`
	IncorrectLoginTries *int       `json:"-"`
	NextLoginTime       *time.Time `json:"-"`
}

type CreateUserRequest struct {
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	OfficeId  int       `json:"officeId,omitempty"`
	Birthdate time.Time `json:"birthdate,omitempty"`
	Password  string    `json:"password,omitempty"`
}

type CreateUserResponse struct {
	User *User `json:"user"`
}

type GetUserRequest struct {
	UserId int `json:"-"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type GetUsersRequest struct {
	OfficeId *int `json:"-"`
}

type GetUsersResponse struct {
	Users []*User `json:"users"`
}

type UpdateUserRequest struct {
	UserId    int     `json:"-"`
	RoleId    *int    `json:"roleId,omitempty"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	OfficeId  *int    `json:"officeId,omitempty"`
}
