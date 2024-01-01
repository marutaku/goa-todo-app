package taskapi

import (
	auth "backend/gen/auth"
	"context"
	"log"
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct {
	logger *log.Logger
}

// NewAuth returns the auth service implementation.
func NewAuth(logger *log.Logger) auth.Service {
	return &authsrvc{logger}
}

// Login to the system
func (s *authsrvc) Login(ctx context.Context, p *auth.LoginPayload) (res *auth.LoginResult, err error) {
	res = &auth.LoginResult{}
	s.logger.Print("auth.login")
	return
}

// Register a new user
func (s *authsrvc) Register(ctx context.Context, p *auth.RegisterPayload) (res *auth.RegisterResult, err error) {
	res = &auth.RegisterResult{}
	s.logger.Print("auth.register")
	return
}

// Logout of the system
func (s *authsrvc) Logout(ctx context.Context, p *auth.LogoutPayload) (err error) {
	s.logger.Print("auth.logout")
	return
}
