package controllers

import (
	"net/http"

	"chat-backend/database"
	"chat-backend/dto"
	"chat-backend/models"
	"chat-backend/service"

	"github.com/gin-gonic/gin"

	"fmt"
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

	// Leer el JSON del body y convertirlo a la struct Usuario
	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(400, gin.H{"error": "JSON malformado: " + err.Error()})
		return
	}

	// Verifico que el email no exista
	err := service.VerificarEmail(database.DB, nuevoUsuario.Email)
	if err != nil {
		fmt.Println("Error:", err.Error())
		c.JSON(400, gin.H{"error": "Error: " + err.Error()}) // Se detiene la creacion y se indica error
		return
	} else {
		fmt.Println("Correo valido")
	}

	// Verifico que la contrasena sea valida
	err = service.ValidarContrasena(nuevoUsuario.Contrasena)
	if err != nil {
		fmt.Println("Error:", err.Error())
		c.JSON(400, gin.H{"error": "Error: " + err.Error()}) // Se detiene la creacion y se indica error
		return
	} else {
		fmt.Println("Contraseña valida")
	}

	// Guardar en la base de datos
	result := database.DB.Create(&nuevoUsuario)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al guardar en base de datos: " + result.Error.Error()})
		return
	}

	// Enviar el usuario recién creado como respuesta
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

func Login(c *gin.Context) {
	var loginDTO dto.LoginRequest

	// Leer JSON y convertirlo a LoginRequest
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de login inválidos: " + err.Error()})
		return
	}

	// Autenticar el usuario (Se enlaza con el service de usuario)
	usuario, err := service.AutenticarUsuario(database.DB, loginDTO.Email, loginDTO.Contrasena)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Enviar respuesta con usuario autenticado
	c.JSON(http.StatusOK, gin.H{
		"usuario": usuario,
	})
}

func BuscarNombreUsuario(c *gin.Context) {
	nombre := c.Param("nombre")
	patron := "%" + nombre + "%" // busca el nombre en cualquier parte del string

	var buscarUsuario []models.Usuario

	result := database.DB.
		Where("LOWER(nombre) LIKE LOWER(?)", patron). // LOWER(username) LIKE LOWER(?) fuerza ambos lados a minúsculas para que no importe si el nombre tiene mayúsculas.
		Find(&buscarUsuario)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "No se encontraron usuarios"})
		return
	}

	c.JSON(200, buscarUsuario)
}

func BuscarIdUsuario(c *gin.Context) {
	id := c.Param("id")

	var buscarUsuario models.Usuario

	result := database.DB.First(&buscarUsuario, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al consultar la base de datos"})
		return
	}

	c.JSON(200, buscarUsuario)
}
