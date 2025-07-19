package models

type Usuario struct {
	ID         int    `gorm:"column:id;primaryKey"`
	Nombre     string `gorm:"column:nombre"`
	Email      string `gorm:"column:email"`
	Contrasena string `gorm:"column:contrasena"`
	Rol        int    `gorm:"column:rol"`
}
