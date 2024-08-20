package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tejasvi541/Go-Server/utils"
)

func Authenticate(context *gin.Context) {
	// Get the token from the header
	token := context.Request.Header.Get("Authorization")

	// If the token is empty
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		context.Abort()
		return
	}

	// Validate the token
	userId, err := utils.ValidateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		context.Abort()
		return
	}

	// Set the user ID in the context
	context.Set("userID", userId)

	// Continue with the request
	context.Next()

}