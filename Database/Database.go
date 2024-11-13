package database

import (
	configs "backend/Configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(configs.DBURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
