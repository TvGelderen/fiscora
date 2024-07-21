package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tvgelderen/budget-buddy/config"
)

const accessTokenKey = "AccessToken"

type CustomClaims struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	jwt.RegisteredClaims
}

func CreateToken(id uuid.UUID, name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Id:    id,
		Name:  name,
		Email: email,
	})

	return token.SignedString([]byte(config.Envs.HMAC))
}

func SetToken(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     accessTokenKey,
		Value:    token,
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func GetId(r *http.Request) (uuid.UUID, error) {
	token, err := getToken(r)
	if err != nil {
		return uuid.UUID{}, err
	}

	parsedToken, err := parseToken(token)
	if err != nil {
		return uuid.UUID{}, err
	}

	return parsedToken.Claims.(*CustomClaims).Id, nil
}

func DeleteToken(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     accessTokenKey,
		Value:    "",
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func getToken(r *http.Request) (string, error) {
	bearer := r.Header.Get("Authorization")
	if bearer != "" {
		splitBearer := strings.Split(bearer, " ")
		return splitBearer[len(splitBearer)-1], nil
	}

	cookie, err := r.Cookie(accessTokenKey)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func parseToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Envs.HMAC), nil
	})

	return parsedToken, err
}
