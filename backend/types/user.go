package types

import "github.com/tvgelderen/fiscora/database"

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func ToUser(user database.User) User {
	return User{
		Username: user.Username,
		Email:    user.Email,
	}
}
