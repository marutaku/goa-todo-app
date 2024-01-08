package repository

import (
	"backend/domain"
	"context"
)

type UserRepositoryInterface interface {
	FindById(ctx context.Context, id domain.UserId) (*domain.User, error)
	FindByName(ctx context.Context, name string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}
