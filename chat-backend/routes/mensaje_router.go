package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func MensajeRouter(c *gin.RouterGroup) {

	c.GET("/mensaje", controllers.GetMensajes)
	c.POST("/mensaje", controllers.CrearMensaje)
	c.PUT("/mensaje/:id", controllers.ActualizarMensaje)
	c.DELETE("/mensaje/:id", controllers.EliminarMensaje)

}
