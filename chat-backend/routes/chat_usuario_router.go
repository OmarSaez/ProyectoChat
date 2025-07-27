package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ChatUsuarioRouter(r *gin.RouterGroup) {

	r.GET("/chat-usuario", controllers.GetChatUsuarios)
	r.POST("/chat-usuario", controllers.CrearChatUsuario)
	r.DELETE("/chat-usuario/:id_usuario/:id_chat", controllers.EliminarChatUsuario)
}
