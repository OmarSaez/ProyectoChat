package controllers

import (
	"chat-backend/database"
	"chat-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener todas las relaciones chat-usuario
func GetChatUsuarios(c *gin.Context) {
	var relaciones []models.ChatUsuario

	result := database.DB.Find(&relaciones)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, relaciones)
}

// Crear una nueva relación chat-usuario
func CrearChatUsuario(c *gin.Context) {
	var nuevaRelacion models.ChatUsuario

	if err := c.ShouldBindJSON(&nuevaRelacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	result := database.DB.Create(&nuevaRelacion)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, nuevaRelacion)
}

// Eliminar una relación chat-usuario
func EliminarChatUsuario(c *gin.Context) {
	idUsuarioParam := c.Param("id_usuario")
	idChatParam := c.Param("id_chat")

	idUsuario, err := strconv.Atoi(idUsuarioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID_Usuario inválido"})
		return
	}
	idChat, err := strconv.Atoi(idChatParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID_Chat inválido"})
		return
	}

	// Crear una instancia para identificar la relación a eliminar
	relacion := models.ChatUsuario{
		IDUsuario: idUsuario,
		IDChat:    idChat,
	}

	// Verificar si existe antes de eliminar
	if err := database.DB.First(&relacion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relación no encontrada"})
		return
	}

	// Eliminar la relación
	if err := database.DB.Delete(&relacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la relación: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Relación eliminada correctamente"})
}
