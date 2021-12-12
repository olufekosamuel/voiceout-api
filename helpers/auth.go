package helpers

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email

	t, err := token.SignedString([]byte("secretstring"))

	if err != nil {
		return "", err
	}

	return t, nil
}
