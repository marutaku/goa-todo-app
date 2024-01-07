package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		repo repository.UserRepositoryInterface
	}
	JWTClaims struct {
		Sub uint32
		jwt.RegisteredClaims
	}
)

func encodeJWTToken(user *domain.User) (string, error) {
	key := os.Getenv("ENCODED_SECRET_KEY")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": user.ID,
		})
	token, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func decodeJWTToken(tokenString string) (uint32, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	} else if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims.Sub, nil
	} else {
		return 0, errors.New("invalid token")
	}
}

func hashPassword(password string) string {
	hashedPassword := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashedPassword[:])
}

func NewAuthInteractor(repo repository.UserRepositoryInterface) *authInteractor {
	return &authInteractor{
		repo: repo,
	}
}

func (u *authInteractor) Login(ctx context.Context, name string, password string) (string, error) {
	hashedPassword := hashPassword(password)
	user, err := u.repo.Find(ctx, name, hashedPassword)
	if err != nil {
		return "", err
	}
	return encodeJWTToken(user)
}

func (u *authInteractor) Register(ctx context.Context, params UserCreateParams) (string, error) {
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
	return encodeJWTToken(createdUser)
}
