package models

type GrupoMiembro struct {
	IDUsuario int  `gorm:"column:id_usuario;primaryKey" json:"id_usuario"`
	IDGrupo   int  `gorm:"column:id_grupo;primaryKey" json:"id_grupo"`
	Admin     bool `gorm:"column:admin"`
}

func (GrupoMiembro) TableName() string {
	return "grupomiembros"
}
