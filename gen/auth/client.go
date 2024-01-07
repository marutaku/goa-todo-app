// Code generated by goa v3.14.1, DO NOT EDIT.
//
// auth client
//
// Command:
// $ goa gen backend/design

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "auth" service client.
type Client struct {
	LoginEndpoint    goa.Endpoint
	RegisterEndpoint goa.Endpoint
}

// NewClient initializes a "auth" service client given the endpoints.
func NewClient(login, register goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint:    login,
		RegisterEndpoint: register,
	}
}

// Login calls the "login" endpoint of the "auth" service.
// Login may return the following errors:
//   - "login_failed" (type LoginFailed)
//   - error: internal error
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *LoginResult, err error) {
	var ires any
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginResult), nil
}

// Register calls the "register" endpoint of the "auth" service.
// Register may return the following errors:
//   - "register_failed" (type RegisterFailed)
//   - error: internal error
func (c *Client) Register(ctx context.Context, p *RegisterPayload) (res *RegisterResult, err error) {
	var ires any
	ires, err = c.RegisterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*RegisterResult), nil
}
