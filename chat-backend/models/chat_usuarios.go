package models

type ChatUsuario struct {
	IDUsuario int `gorm:"column:id_usuario;primaryKey" json:"id_usuario"`
	IDChat    int `gorm:"column:id_chat;primaryKey" json:"id_chat"`
}

func (ChatUsuario) TableName() string {
	return "chatusuarios"
}
