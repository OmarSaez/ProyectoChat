package controllers

import (
	"net/http"

	"chat-backend/database"
	"chat-backend/models"

	"github.com/gin-gonic/gin"
)

func GetUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
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
