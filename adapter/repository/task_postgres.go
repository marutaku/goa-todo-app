package repository

import (
	"backend/domain"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

type TaskRecord struct {
	// Unique ID
	ID uint32 `gorm:"primaryKey"`
	// Name of the todo
	Name string `gorm:"not null"`
	// Description of the todo
	Description string `gorm:"not null;default ''"`
	// Whether or not the todo is done
	Done bool `gorm:"not null;default false"`
	// When the todo was done in ISO format
	DoneAt string `gorm:"default ''"`
	// Who did the todo
	DoneBy string `gorm:"default ''"`
	// When the todo was created in ISO format
	CreatedAt string `gorm:"not null"`
	// Who created the todo
	CreatedBy string `gorm:"not null"`
}

func (t *TaskRecord) ToDomain() (*domain.Task, error) {
	var doneAt *time.Time
	if t.DoneAt != "" {
		parsedDate, err := time.Parse(time.RFC3339, t.DoneAt)
		if err != nil {
			return nil, err
		}
		doneAt = &parsedDate
	} else {
		doneAt = nil
	}
	createdAt, err := time.Parse(time.RFC3339, t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return domain.NewTask(
		domain.TaskId(t.ID),
		t.Name,
		t.Description,
		t.Done,
		doneAt,
		t.DoneBy,
		createdAt,
		t.CreatedBy,
	)
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	db.AutoMigrate(&TaskRecord{})
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) FindAll(ctx context.Context, name *string, done *bool, createdBy *string) ([]*domain.Task, error) {
	var taskRecords []*TaskRecord
	criteria := &TaskRecord{}
	if name != nil {
		criteria.Name = *name
	}
	if done != nil {
		criteria.Done = *done
	}
	if createdBy != nil {
		criteria.CreatedBy = *createdBy
	}
	if err := t.db.WithContext(ctx).Where(criteria).Find(&taskRecords).Error; err != nil {
		return nil, err
	}
	var tasks []*domain.Task
	for _, task := range taskRecords {
		domainTask, err := task.ToDomain()
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, domainTask)
	}
	return tasks, nil
}

func (t *TaskRepository) FindOne(ctx context.Context, id domain.TaskId) (*domain.Task, error) {
	var taskRecord TaskRecord
	if err := t.db.WithContext(ctx).First(&taskRecord, id.UInt32()).Error; err != nil {
		return nil, err
	}
	return taskRecord.ToDomain()
}

func (t *TaskRepository) Create(ctx context.Context, newTask *domain.Task) (*domain.Task, error) {
	taskRecord := &TaskRecord{
		Name:        newTask.Name,
		Description: newTask.Description,
		Done:        false,
		CreatedAt:   newTask.CreatedAt.Format(time.RFC3339),
	}
	if err := t.db.WithContext(ctx).Create(&taskRecord).Error; err != nil {
		return nil, err
	}
	fmt.Println(taskRecord)
	return taskRecord.ToDomain()
}

func (t *TaskRepository) Update(ctx context.Context, newTask *domain.Task) (*domain.Task, error) {
	var taskRecord *TaskRecord
	if err := t.db.WithContext(ctx).First(&taskRecord, newTask.ID).Error; err != nil {
		return nil, err
	}
	taskRecord.Name = newTask.Name
	taskRecord.Description = newTask.Description
	if err := t.db.WithContext(ctx).Save(&taskRecord).Error; err != nil {
		return nil, err
	}
	return taskRecord.ToDomain()
}

func (t *TaskRepository) Delete(ctx context.Context, id domain.TaskId) error {
	if err := t.db.WithContext(ctx).Delete(&TaskRecord{}, id.UInt32()).Error; err != nil {
		return err
	}
	return nil
}
