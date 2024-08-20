package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken( email, userId string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	if err != nil {
		return "", err
	}

	return token.SignedString([]byte("Lolllll"))
}