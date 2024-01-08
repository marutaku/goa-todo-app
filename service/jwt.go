package service

import (
	"backend/domain"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthService struct {
	secretKey string
}

type JWTClaims struct {
	Sub string `json:"sub"`
	jwt.RegisteredClaims
}

func NewJwTAuthService() *JWTAuthService {
	return &JWTAuthService{
		secretKey: os.Getenv("ENCODED_SECRET_KEY"),
	}
}

func (s *JWTAuthService) EncodeJWTToken(user *domain.User) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": strconv.FormatUint(uint64(user.ID), 10),
		})
	token, err = t.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *JWTAuthService) DecodeJWTToken(tokenString string) (uint32, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return 0, &domain.AuthError{Err: errors.New("invalid token")}
	} else if claims, ok := token.Claims.(*JWTClaims); ok {
		userId, err := strconv.ParseUint(claims.Sub, 10, 32)
		if err != nil {
			return 0, err
		}

		return uint32(userId), nil
	} else {
		fmt.Println(err)
		return 0, &domain.AuthError{Err: errors.New("invalid token")}
	}
}

func (s *JWTAuthService) VerifyToken(tokenString string) (uint32, error) {
	userId, err := s.DecodeJWTToken(tokenString)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
