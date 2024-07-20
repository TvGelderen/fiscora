package auth

import (
	"github.com/tvgelderen/budget-buddy/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {
	GoogleConfig *oauth2.Config
}

func NewAuthService() *AuthService {
	env := config.Envs

	googleConfig := &oauth2.Config{
		ClientID:     env.GoogleID,
		ClientSecret: env.GoogleSecret,
		RedirectURL:  env.GoogleCallback,
		Endpoint:     google.Endpoint,
	}

	return &AuthService{
		GoogleConfig: googleConfig,
	}
}

type GoogleUser struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"piture"`
	Locale        string `json:"locale"`
}
