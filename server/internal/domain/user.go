package domain

import "time"

type User struct {
	Id                  *int       `json:"id,omitempty" example:"44"`
	RoleId              *int       `json:"roleId,omitempty" example:"1"`
	Email               *string    `json:"email,omitempty" example:"something@mail.com"`
	Password            *string    `json:"-"`
	FirstName           *string    `json:"firstName,omitempty" example:"Alex"`
	LastName            *string    `json:"lastName,omitempty" example:"Herbarien"`
	OfficeId            *int       `json:"officeId,omitempty" example:"20"`
	Birthdate           *string    `json:"birthdate,omitempty" example:"2017-12-31"`
	Age                 *int       `json:"age,omitempty" example:"30"`
	Active              *bool      `json:"active,omitempty"`
	IncorrectLoginTries *int       `json:"-"`
	NextLoginTime       *time.Time `json:"-"`
}

type CreateUserRequest struct {
	Email     string `json:"email" example:"something@mail.com"`
	FirstName string `json:"firstName" example:"Alex"`
	LastName  string `json:"lastName" example:"Herbarien"`
	OfficeId  int    `json:"officeId" example:"20"`
	Birthdate string `json:"birthdate" example:"2017-12-31"`
	Password  string `json:"password" example:"renatikadik22"`
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
	RoleId    *int    `json:"roleId,omitempty" example:"1"`
	Email     *string `json:"email,omitempty" example:"something@mail.com"`
	FirstName *string `json:"firstName,omitempty" example:"Alex"`
	LastName  *string `json:"lastName,omitempty" example:"Herbarien"`
	OfficeId  *int    `json:"officeId,omitempty" example:"20"`
}

type UserLogin struct {
	TimeSpent  *string `json:"timeSpent" example:"17:00:00"`
	LoginTime  string  `json:"loginTime" example:"2017-12-31 17:00"`
	LogoutTime *string `json:"logoutTime" example:"2017-12-31 17:00"`
	Error      *string `json:"error" example:"New login detected without logout"`
}

type GetUserLoginsResponse struct {
	UserLogins             []UserLogin `json:"userLogins"`
	LastLoginErrorDatetime *string     `json:"lastLoginErrorDatetime" example:"2017-12-31 17:00"`
	NumberOfCrashes        int         `json:"numberOfCrashes" example:"3"`
}
