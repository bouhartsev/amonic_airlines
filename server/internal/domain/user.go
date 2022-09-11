package domain

import "time"

type User struct {
	Id        *int       `json:"id,omitempty"`
	RoleId    *int       `json:"roleId,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Password  *string    `json:"password,omitempty"`
	FirstName *string    `json:"firstName,omitempty"`
	LastName  *string    `json:"lastName,omitempty"`
	OfficeId  *int       `json:"officeId,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Active    *bool      `json:"active,omitempty"`
}
