package controllers

import (
	"chat-backend/database"
	"chat-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGrupos(c *gin.Context) {

	var grupos []models.Grupo

	result := database.DB.Find(&grupos)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, grupos)
}

func CrearGrupo(c *gin.Context) {

	var nuevoGrupo models.Grupo

	//Mapea lo recibido de la peticion a la structura del grupo
	if err := c.ShouldBindJSON(&nuevoGrupo); err != nil {
		c.JSON(400, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	result := database.DB.Create(&nuevoGrupo)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, nuevoGrupo)
}

func ActualizarGrupo(c *gin.Context) {

	// Obtener el ID desde la URL
	idParam := c.Param("id")

	// Trasnformar el ID a un entero (la URL viene en string)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	// Buscar al usuario en la base de datos
	var grupo models.Grupo
	if err := database.DB.First(&grupo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Leer el JSON con los nuevos datos
	if err := c.ShouldBindJSON(&grupo); err != nil {
		c.JSON(400, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	// Guardar los cambios en la base de datos
	if err := database.DB.Save(&grupo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar usuario: " + err.Error()})
		return
	}

	c.JSON(200, grupo)

}

func EliminarGrupo(c *gin.Context) {

	idParam := c.Param("id")

	var grupo models.Grupo

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	if err := database.DB.First(&grupo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Grupo no encontrado"})
		return
	}

	result := database.DB.Delete(&grupo)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"mensaje": "Grupo eliminado correctamente"})
}
