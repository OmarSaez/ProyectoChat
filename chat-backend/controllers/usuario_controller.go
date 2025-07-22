package controllers

import (
	"net/http"

	"chat-backend/database"
	"chat-backend/models"

	"github.com/gin-gonic/gin"

	"strconv"
)

func GetUsuarios(c *gin.Context) {
	var usuarios []models.Usuario

	// Obtener todos los usuarios mediante una consulta
	result := database.DB.Find(&usuarios)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func CrearUsuario(c *gin.Context) {
	var nuevoUsuario models.Usuario

	// 1. Leer el JSON del body y convertirlo a la struct Usuario
	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(400, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	// 2. Guardar en la base de datos
	result := database.DB.Create(&nuevoUsuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al guardar en base de datos: " + result.Error.Error()})
		return
	}

	// 3. Enviar el usuario recién creado como respuesta
	c.JSON(201, nuevoUsuario)
}

func ActualizarUsuario(c *gin.Context) {
	// Obtener el ID desde la URL
	idParam := c.Param("id")

	// Trasnformar el ID a un entero (la URL viene en string)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	// Buscar al usuario en la base de datos
	var usuario models.Usuario
	if err := database.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Leer el JSON con los nuevos datos
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	// Guardar los cambios en la base de datos
	if err := database.DB.Save(&usuario).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar usuario: " + err.Error()})
		return
	}

	c.JSON(200, usuario)
}

func EliminarUsuario(c *gin.Context) {
	id := c.Param("id")

	// Buscar al usuario por ID
	var usuario models.Usuario
	if err := database.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Eliminar el usuario
	if err := database.DB.Delete(&usuario).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al eliminar usuario: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"mensaje": "Usuario eliminado correctamente"})
}
