package repository

import (
	"backend/domain"
	"context"
)

// FIXME: This is a workaround for use struct in gorm where clause
type TaskCriteria map[string]interface{}

type TaskRepositoryInterface interface {
	FindAll(ctx context.Context, criteria TaskCriteria) ([]*domain.Task, error)
	FindOne(ctx context.Context, id domain.TaskId) (*domain.Task, error)
	Create(ctx context.Context, task *domain.Task) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) (*domain.Task, error)
	Delete(ctx context.Context, id domain.TaskId) error
}
