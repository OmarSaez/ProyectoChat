package models

type Contacto struct {
	IDUsuario  int `gorm:"column:id_usuario;primaryKey"`
	IDContacto int `gorm:"column:id_contacto;primaryKey"`
}
