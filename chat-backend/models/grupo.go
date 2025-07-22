package models

type Grupo struct {
	ID     int    `gorm:"column:id;primaryKey"`
	Nombre string `gorm:"column:nombre"`
	IDChat int    `gorm:"column:id_chat;unique" json:"id_chat"`
}
