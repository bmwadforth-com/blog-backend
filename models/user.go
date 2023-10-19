package models

import "time"

// UserModel A model that describes a user
// @Description A model that describes a user
type UserModel struct {
	UserId      string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	CreatedDate time.Time `json:"created"`
	UpdatedDate time.Time `json:"updated"`
	DocumentRef string    `json:"-"`
	Active      bool      `json:"-"`
}

// UserSessionModel A model that describes a user session
// @Description A model that describes a user session
type UserSessionModel struct {
	UserId   string    `json:"id"`
	Username string    `json:"username"`
	LoggedIn time.Time `json:"loggedIn"`
	Token    string    `json:"-"`
	Active   bool      `json:"-"`
}

// CreateUserRequest New user
// @Description New user request
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUserRequest Login user
// @Description Login user request
type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
