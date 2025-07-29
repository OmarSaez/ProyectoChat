package routes

import (
	"chat-backend/controllers"

	"github.com/gin-gonic/gin"
)

func GrupoMiembroRouter(r *gin.RouterGroup) {

	r.GET("/grupo-miembro", controllers.GetGrupoMiembros)
	r.POST("/grupo-miembro", controllers.CrearGrupoMiembro)
	r.PUT("/grupo-miembro/:id_usuario/:id_grupo", controllers.EliminarGrupoMiembro)
	r.DELETE("/grupo-miembro/:id_usuario/:id_grupo", controllers.EliminarGrupoMiembro)
}
