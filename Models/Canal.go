package models

import "gorm.io/gorm"

type Canal struct {
	gorm.Model
	Nombre   string    `json:"nombre"`
	Mensajes []Mensaje `json:"mensajes"`
}
