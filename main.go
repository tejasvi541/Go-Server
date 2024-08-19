package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8080")
}