package main

import (
	"chat-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect() //Me conecto a la base de datos

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // Servidor en localhost:8080
}
