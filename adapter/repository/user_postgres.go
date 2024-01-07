package repository

import (
	"backend/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type AuthRecord struct {
	ID        uint32 `gorm:"primaryKey"`
	Name      string `gorm:"not null; unique"`
	Password  string `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
}

func (r *AuthRecord) ToDomain() (*domain.User, error) {
	createdAt, err := time.Parse(time.RFC3339, r.CreatedAt)
	if err != nil {
		return nil, err
	}
	return domain.NewUser(
		r.ID,
		r.Name,
		r.Password,
		createdAt,
	)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&AuthRecord{})
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	var record AuthRecord
	result := r.db.WithContext(ctx).Where("name = ?", name).First(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	user, err := record.ToDomain()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	record := &AuthRecord{
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
	r.db.WithContext(ctx).Create(&record)
	user, err := record.ToDomain()
	if err != nil {
		return nil, err
	}
	return user, nil
}
