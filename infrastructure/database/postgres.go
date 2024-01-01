package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresHandler struct {
	db *gorm.DB
}

func NewPostgresHandler(config *config) (*PostgresHandler, error) {
	db, err := gorm.Open(postgres.Open(config.dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresHandler{db: db}, nil
}
