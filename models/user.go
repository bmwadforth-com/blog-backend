package models

import "time"

// UserModel A model that describes a user
// @Description A model that describes a user
type UserModel struct {
	UserId      string    `json:"id" firestore:"id"`
	Username    string    `json:"username" firestore:"username"`
	Password    string    `json:"-" firestore:"password"`
	CreatedDate time.Time `json:"created" firestore:"created"`
	UpdatedDate time.Time `json:"updated" firestore:"updated"`
	DocumentRef string    `json:"-" firestore:"-"`
	Active      bool      `json:"-" firestore:"active"`
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

// UserStatusModel User status model
// @Description User status model
type UserStatusModel struct {
	Username      string `json:"username"`
	Active        bool   `json:"active"`
	LoggedInSince string `json:"loggedInSince"`
}
