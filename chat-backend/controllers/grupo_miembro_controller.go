package controllers

import (
	"chat-backend/database"
	"chat-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener todos los grupo-miembros
func GetGrupoMiembros(c *gin.Context) {
	var miembros []models.GrupoMiembro

	result := database.DB.Find(&miembros)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, miembros)
}

// Crear una nueva relación grupo-miembro
func CrearGrupoMiembro(c *gin.Context) {
	var nuevoMiembro models.GrupoMiembro

	if err := c.ShouldBindJSON(&nuevoMiembro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	result := database.DB.Create(&nuevoMiembro)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, nuevoMiembro)
}

// Actualizar una relación grupo-miembro existente
func ActualizarGrupoMiembro(c *gin.Context) {
	idUsuarioParam := c.Param("id_usuario")
	idGrupoParam := c.Param("id_grupo")

	idUsuario, err := strconv.Atoi(idUsuarioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDUsuario inválido"})
		return
	}
	idGrupo, err := strconv.Atoi(idGrupoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDGrupo inválido"})
		return
	}

	var miembro models.GrupoMiembro
	if err := database.DB.
		Where("id_usuario = ? AND id_grupo = ?", idUsuario, idGrupo).
		First(&miembro).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Miembro no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&miembro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	if err := database.DB.Save(&miembro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar miembro: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, miembro)
}

// Eliminar una relación grupo-miembro
func EliminarGrupoMiembro(c *gin.Context) {
	idUsuarioParam := c.Param("id_usuario")
	idGrupoParam := c.Param("id_grupo")

	idUsuario, err := strconv.Atoi(idUsuarioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDUsuario inválido"})
		return
	}
	idGrupo, err := strconv.Atoi(idGrupoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDGrupo inválido"})
		return
	}

	relacion := models.GrupoMiembro{
		IDUsuario: idUsuario,
		IDGrupo:   idGrupo,
	}

	if err := database.DB.Delete(&relacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar relación: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Relación eliminada correctamente"})
}
