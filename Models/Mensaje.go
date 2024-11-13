package models

import "gorm.io/gorm"

type Mensaje struct {
	gorm.Model
	Username  string `json:"username"`
	Contenido string `json:"contenido"`
	CanalID   uint   `json:"canal_id"`
}
