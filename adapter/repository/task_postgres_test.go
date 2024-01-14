package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

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
	if err != nil {
		return nil, nil, err
	}
	err = mockDB.AutoMigrate(&TaskRecord{})
	if err != nil {
		return nil, nil, err
	}
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

func TestFindAll(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Error(err)
	}
	rows := sqlmock.NewRows([]string{"id", "name", "description", "done", "done_at", "done_by", "created_at", "created_by"})
	for i := 1; i <= 10; i++ {
		rows.AddRow(i, fmt.Sprintf("name-%d", i), fmt.Sprintf("description-%d", i), false, nil, nil, time.Now().Format(time.RFC3339), 1)
	}
	logger := log.New(os.Stderr, "[taskapi] ", log.Ltime)
	repo := NewTaskRepository(mockDB, logger)
	if repo == nil {
		t.Error("repo is nil")
	}
	mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "task_records" WHERE "task_records"."created_by" = $1`),
	).WithArgs(
		1,
	).WillReturnRows(
		rows,
	)
	var mockString *string
	tasks, err := repo.FindAll(context.Background(), mockString, nil, 1)
	if err != nil {
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	if len(tasks) != 10 {
		t.Errorf("expected 10 tasks, got %d", len(tasks))
	}
}
