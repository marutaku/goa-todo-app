package presenter

import (
	authService "backend/gen/auth"
)

type AuthPresenter struct{}

func NewAuthPresenter() *AuthPresenter {
	return &AuthPresenter{}
}

func (p *AuthPresenter) LoginOutput(token string) *authService.LoginResult {
	res := &authService.LoginResult{Token: token}
	return res
}

func (p *AuthPresenter) RegisterOutput(token string) *authService.RegisterResult {
	res := &authService.RegisterResult{Token: token}
	return res
}
