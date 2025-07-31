package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ContactoRouter(r *gin.RouterGroup) {

	r.GET("/contacto", controllers.GetContactos)
	r.POST("/contacto", controllers.CrearContacto)
	r.DELETE("/contacto/:id", controllers.EliminarContacto)
}
