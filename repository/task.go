package repository

import (
	"backend/domain"
	"context"
)

type TaskRepositoryInterface interface {
	FindAll(ctx context.Context) ([]*domain.Task, error)
	FindOne(ctx context.Context, id domain.TaskId) (*domain.Task, error)
	Create(ctx context.Context, task *domain.Task) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) (*domain.Task, error)
	Delete(ctx context.Context, id domain.TaskId) error
}
