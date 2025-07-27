package main

import (
	"chat-backend/database"

	"github.com/gin-gonic/gin"

	"chat-backend/routes"
)

func main() {
	database.Connect() //Me conecto a la base de datos

	r := gin.Default()

	//testeo
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Crear router con grupo de rutas
	api := r.Group("/api") //Agrupamos todas las URL con /api

	routes.UserRouter(api)
	routes.ChatRouter(api)
	routes.GrupoRouter(api)
	routes.MensajeRouter(api)
	routes.ChatUsuarioRouter(api)

	r.Run(":8080") // Servidor en localhost:8080
}
