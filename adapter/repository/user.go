package repository

import (
	"backend/domain"
	"context"
)

type UserRepositoryInterface interface {
	Find(ctx context.Context, name string, password string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}
