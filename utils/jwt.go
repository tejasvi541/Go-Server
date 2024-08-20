package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)



func GenerateToken( email string, userId int64) (string, error) {
	godotenv.Load()
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenString string) (error) {	
	godotenv.Load()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return  err
	}

	tokenIsValid := token.Valid
	if !tokenIsValid {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return err
	}
	return nil
}