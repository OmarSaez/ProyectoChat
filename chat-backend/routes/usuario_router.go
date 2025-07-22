package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/usuarios", controllers.GetUsuarios)
	r.POST("/usuarios", controllers.CrearUsuario)
	r.PUT("/usuarios/:id", controllers.ActualizarUsuario)
	r.DELETE("/usuarios/:id", controllers.EliminarUsuario)
	r.POST("/usuarios/login", controllers.Login)

	return r
}
