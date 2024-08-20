package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tejasvi541/Go-Server/src/models"
	"github.com/tejasvi541/Go-Server/utils"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)

}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Authenticate(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token , err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}