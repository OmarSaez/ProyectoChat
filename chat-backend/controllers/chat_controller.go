package controllers

import (
	"chat-backend/database"
	"chat-backend/models"

	"github.com/gin-gonic/gin"
)

func GetChat(c *gin.Context) {
	var chats []models.Chat

	resultado := database.DB.Find(&chats)

	if resultado.Error != nil {
		c.JSON(404, gin.H{"error": resultado.Error.Error()})
		return
	}

	c.JSON(200, chats)
}

func CrearChat(c *gin.Context) {

	var nuevoChat models.Chat

	result := database.DB.Create(&nuevoChat)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, nuevoChat)
}

func EliminarChat(c *gin.Context) {

	id := c.Param("id")

	var chat models.Chat
	if err := database.DB.First(&chat, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Chat no encontrado"})
		return
	}

	result := database.DB.Delete(&chat)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"mensaje": "El chat se elimino correctamente"})
}
