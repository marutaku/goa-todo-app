package models

import (
	task "backend/gen/task"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDatabase(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&task.BackendStoredTask{})
	db.Logger = db.Logger.LogMode(logger.Info)
	return db, nil
}
