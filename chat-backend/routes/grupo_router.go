package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func GrupoRouter(r *gin.RouterGroup) {

	r.GET("/grupo", controllers.GetGrupos)
	r.POST("/grupo", controllers.CrearGrupo)
	r.PUT("/grupo/:id", controllers.ActualizarGrupo)
	r.DELETE("/grupo/:id", controllers.EliminarGrupo)
}
