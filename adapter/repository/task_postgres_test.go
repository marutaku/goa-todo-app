package repository

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestNewTaskRepository(t *testing.T) {
	mockDB, _, err := NewDbMock()
	if err != nil {
		t.Error(err)
	}
	logger := log.New(os.Stderr, "[taskapi] ", log.Ltime)
	repo := NewTaskRepository(mockDB, logger)
	if repo == nil {
		t.Error("repo is nil")
	}
}
