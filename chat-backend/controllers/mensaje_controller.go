package controllers

import (
	"chat-backend/database"
	"chat-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener todos los mensajes
func GetMensajes(c *gin.Context) {
	var mensajes []models.Mensaje

	result := database.DB.Find(&mensajes)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, mensajes)
}

// Crear un nuevo mensaje
func CrearMensaje(c *gin.Context) {
	var nuevoMensaje models.Mensaje

	if err := c.ShouldBindJSON(&nuevoMensaje); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	result := database.DB.Create(&nuevoMensaje)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, nuevoMensaje)
}

// Actualizar un mensaje existente
func ActualizarMensaje(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var mensaje models.Mensaje
	if err := database.DB.First(&mensaje, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mensaje no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&mensaje); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	if err := database.DB.Save(&mensaje).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar mensaje: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, mensaje)
}

// Eliminar un mensaje
func EliminarMensaje(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var mensaje models.Mensaje
	if err := database.DB.First(&mensaje, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mensaje no encontrado"})
		return
	}

	result := database.DB.Delete(&mensaje)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Mensaje eliminado correctamente"})
}
