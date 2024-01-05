package repository

import (
	"backend/domain"
	"context"

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

type TaskCriteria TaskRecord

func (t *TaskRecord) ToDomain() (*domain.Task, error) {
	return domain.NewTask(
		domain.TaskId(t.ID),
		t.Name,
		t.Description,
		t.Done,
		t.DoneAt,
		t.DoneBy,
		t.CreatedAt,
		t.CreatedBy,
	)
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	db.AutoMigrate(&TaskRecord{})
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) FindAll(ctx context.Context, taskCondition TaskCriteria) ([]*domain.Task, error) {
	criteria := map[string]interface{}{}
	if taskCondition.CreatedBy != "" {
		criteria["created_by"] = taskCondition.CreatedBy
	}
	if taskCondition.Name != "" {
		criteria["name"] = taskCondition.Name
	}
	var taskRecords []*TaskRecord
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

func (t *TaskRepository) Update(ctx context.Context, id domain.TaskId, name string, description string) (*domain.Task, error) {
	var taskRecord TaskRecord
	if err := t.db.WithContext(ctx).First(&taskRecord, id.UInt32()).Error; err != nil {
		return nil, err
	}
	taskRecord.Name = name
	taskRecord.Description = description
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
