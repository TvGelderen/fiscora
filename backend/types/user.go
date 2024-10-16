package types

import "github.com/tvgelderen/fiscora/repository"

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func ToUser(user *repository.User) User {
	return User{
		Username: user.Username,
		Email:    user.Email,
	}
}
