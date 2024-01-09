package controller

import (
	"backend/adapter/presenter"
	"backend/adapter/repository"
	"backend/domain"
	authService "backend/gen/auth"
	"backend/infrastructure/database"
	"backend/usecase"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type authController struct {
	logger    *log.Logger
	usecase   usecase.AuthUseCase
	presenter *presenter.AuthPresenter
}

func NewAuthController(logger *log.Logger) *authController {
	config := database.NewPostgresConfig()
	db, err := database.NewPostgresDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewUserRepository(db, logger)
	return &authController{
		logger:    logger,
		usecase:   usecase.NewAuthInteractor(repo),
		presenter: presenter.NewAuthPresenter(),
	}
}

func (c *authController) Login(ctx context.Context, params *authService.LoginPayload) (*authService.LoginResult, error) {
	c.logger.Print("auth.login")
	token, err := c.usecase.Login(ctx, params.Username, params.Password)
	if err != nil {
		var authErr *domain.AuthError
		if err == gorm.ErrRecordNotFound {
			c.logger.Print("auth.login failed with ErrRecordNotFound")
			return nil, &authService.AuthFlowFailed{Name: "login_failed", Message: "username or password is incorrect"}
		} else if errors.As(err, &authErr) {
			c.logger.Print("auth.login failed with AuthError")
			return nil, &authService.AuthFlowFailed{Name: "login_failed", Message: "username or password is incorrect"}
		}
		c.logger.Print("auth.login failed with unknown error")
		return nil, err
	}
	return c.presenter.LoginOutput(token), nil
}

func (c *authController) Register(ctx context.Context, params *authService.RegisterPayload) (*authService.RegisterResult, error) {
	c.logger.Print("auth.register")
	token, err := c.usecase.Register(ctx, usecase.UserCreateParams{
		Name:     params.Username,
		Password: params.Password,
	})
	if err != nil {
		var authErr *domain.AuthError
		if errors.As(err, &authErr) {
			return nil, &authService.AuthFlowFailed{Name: "register_failed", Message: authErr.Error()}

		}
	}
	return c.presenter.RegisterOutput(token), nil
}
