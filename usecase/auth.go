package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"backend/service"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type (
	AuthUseCase interface {
		Login(ctx context.Context, name string, password string) (string, error)
		Register(ctx context.Context, params UserCreateParams) (string, error)
	}
	UserCreateParams struct {
		Name     string
		Password string
	}
	authInteractor struct {
		repo       repository.UserRepositoryInterface
		jwtService *service.JWTAuthService
	}
)

func hashPassword(password string) string {
	hashedPassword := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashedPassword[:])
}

func NewAuthInteractor(repo repository.UserRepositoryInterface) *authInteractor {
	return &authInteractor{
		repo:       repo,
		jwtService: service.NewJwTAuthService(repo),
	}
}

func (u *authInteractor) Login(ctx context.Context, name string, password string) (string, error) {
	hashedPassword := hashPassword(password)
	user, err := u.repo.FindByName(ctx, name)
	if err != nil {
		return "", err
	} else if user == nil {
		return "", &domain.AuthError{Err: errors.New("login failed")}
	}
	if user.Password != hashedPassword {
		return "", &domain.AuthError{Err: errors.New("login failed")}
	}
	return u.jwtService.EncodeJWTToken(user)
}

func (u *authInteractor) Register(ctx context.Context, params UserCreateParams) (string, error) {
	existingUser, err := u.repo.FindByName(ctx, params.Name)
	if err != nil {
		return "", err
	}
	if existingUser != nil {
		return "", &domain.AuthError{Err: errors.New("user already exists")}
	}
	hashedPassword := hashPassword(params.Password)
	user, err := domain.NewUser(
		0,
		params.Name,
		hashedPassword,
		time.Now(),
	)
	if err != nil {
		return "", err
	}
	createdUser, err := u.repo.Create(ctx, user)
	if err != nil {
		return "", err
	}
	return u.jwtService.EncodeJWTToken(createdUser)
}
