package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ChatRouter(r *gin.RouterGroup) {

	r.GET("/chat", controllers.GetChat)
	r.POST("/chat", controllers.CrearChat)
	r.DELETE("/chat/:id", controllers.EliminarChat)
}
