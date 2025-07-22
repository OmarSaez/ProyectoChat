package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	r.GET("/usuario", controllers.GetUsuarios)
	r.POST("/usuario", controllers.CrearUsuario)
	r.PUT("/usuario/:id", controllers.ActualizarUsuario)
	r.DELETE("/usuario/:id", controllers.EliminarUsuario)
	r.POST("/usuario/login", controllers.Login)
	r.POST("/usuario/:nombre", controllers.BuscarNombre)
}
