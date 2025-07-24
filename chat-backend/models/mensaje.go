package models

import "time"

type Mensaje struct {
	ID         int        `gorm:"column:id;primaryKey"`
	IDUsuario  int        `gorm:"column:id_usuario" json:"id_usuario"`
	IDChat     int        `gorm:"column:id_chat" json:"id_chat"`
	Contenido  string     `gorm:"column:contenido"`
	FechaEnvio *time.Time `gorm:"column:fechaenvio; default:CURRENT_TIMESTAMP"` //Puntero para que no haga un insert y la fecha la ponga la BD
}
