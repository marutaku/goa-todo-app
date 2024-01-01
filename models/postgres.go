package models

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(config *config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Task{})
	return db, nil
}
