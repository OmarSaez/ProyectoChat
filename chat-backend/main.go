package main

import (
	"chat-backend/database"

	"github.com/gin-gonic/gin"

	"chat-backend/routes"
)

func main() {
	database.Connect() //Me conecto a la base de datos

	d := gin.Default()

	//testeo
	d.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Crear router con rutas ya cargadas desde routes/
	r := routes.SetupRouter()

	r.Run(":8080") // Servidor en localhost:8080
}
