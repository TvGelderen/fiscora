package types

import "github.com/tvgelderen/budget-buddy/database"

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
