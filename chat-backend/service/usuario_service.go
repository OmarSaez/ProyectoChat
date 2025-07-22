package service

import (
	"chat-backend/models" // ajusta el import según tu estructura de proyecto
	"errors"

	"gorm.io/gorm"
)

// Creando un metodo para "hacer mas segura" una contraseña
func ValidarContrasena(contrasena string) error {
	if len(contrasena) < 5 {
		return errors.New("la contrasena debe tener al menos 5 caracteres")
	}

	return nil //si esta bien de vuelvo vacio
}

// Verifica si el email ya existe (para usuarios nuevas)
func VerificarEmail(db *gorm.DB, email string) error {
	var usuario models.Usuario

	//Consulta para buscar el correo
	err := db.Where("email = ?", email).First(&usuario).Error

	//Si no dio error significa que encontro el correo
	if err == nil {
		return errors.New("el email ya existe")
	}

	// Si da error por no encontrarlo, entonces esta libre el correo para usar
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return err // Otro error inesperado
}

// Autenticación simple
func AutenticarUsuario(db *gorm.DB, email string, contrasena string) (*models.Usuario, error) {
	var usuario models.Usuario

	// Buscar usuario por email
	if err := db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	// Comparar contraseñas en texto plano
	if usuario.Contrasena != contrasena {
		return nil, errors.New("contraseña incorrecta")
	}

	return &usuario, nil
}
