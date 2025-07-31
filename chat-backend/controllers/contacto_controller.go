package controllers

import (
	"chat-backend/database"
	"chat-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener todos los contactos
func GetContactos(c *gin.Context) {
	var contactos []models.Contacto

	result := database.DB.Find(&contactos)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, contactos)
}

// Crear un nuevo contacto
func CrearContacto(c *gin.Context) {
	var nuevoContacto models.Contacto

	if err := c.ShouldBindJSON(&nuevoContacto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	result := database.DB.Create(&nuevoContacto)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, nuevoContacto)
}

// Actualizar un contacto existente (pendiente evaluar si hace falta)
func ActualizarContacto(c *gin.Context) {
	idUsuarioParam := c.Param("id_usuario")
	idContactoParam := c.Param("id_contacto")

	idUsuario, err := strconv.Atoi(idUsuarioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDUsuario inválido"})
		return
	}
	idContacto, err := strconv.Atoi(idContactoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDContacto inválido"})
		return
	}

	var contacto models.Contacto
	if err := database.DB.
		Where("id_usuario = ? AND id_contacto = ?", idUsuario, idContacto).
		First(&contacto).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contacto no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&contacto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	if err := database.DB.Save(&contacto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar contacto: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacto)
}

// Eliminar un contacto
func EliminarContacto(c *gin.Context) {
	idUsuarioParam := c.Param("id_usuario")
	idContactoParam := c.Param("id_contacto")

	idUsuario, err := strconv.Atoi(idUsuarioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDUsuario inválido"})
		return
	}
	idContacto, err := strconv.Atoi(idContactoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDContacto inválido"})
		return
	}

	relacion := models.Contacto{
		IDUsuario:  idUsuario,
		IDContacto: idContacto,
	}

	if err := database.DB.Delete(&relacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar relación: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Contacto eliminado correctamente"})
}
