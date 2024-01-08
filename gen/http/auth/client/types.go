// Code generated by goa v3.14.1, DO NOT EDIT.
//
// auth HTTP client types
//
// Command:
// $ goa gen backend/design

package client

import (
	auth "backend/gen/auth"

	goa "goa.design/goa/v3/pkg"
)

// LoginRequestBody is the type of the "auth" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	// Username to login with
	Username string `form:"username" json:"username" xml:"username"`
	// Password to login with
	Password string `form:"password" json:"password" xml:"password"`
}

// RegisterRequestBody is the type of the "auth" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	// Username to register with
	Username string `form:"username" json:"username" xml:"username"`
	// Password to register with
	Password string `form:"password" json:"password" xml:"password"`
}

// LoginResponseBody is the type of the "auth" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	// JWT token to use for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// RegisterResponseBody is the type of the "auth" service "register" endpoint
// HTTP response body.
type RegisterResponseBody struct {
	// JWT token to use for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// LoginLoginFailedResponseBody is the type of the "auth" service "login"
// endpoint HTTP response body for the "login_failed" error.
type LoginLoginFailedResponseBody struct {
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RegisterRegisterFailedResponseBody is the type of the "auth" service
// "register" endpoint HTTP response body for the "register_failed" error.
type RegisterRegisterFailedResponseBody struct {
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// NewLoginRequestBody builds the HTTP request body from the payload of the
// "login" endpoint of the "auth" service.
func NewLoginRequestBody(p *auth.LoginPayload) *LoginRequestBody {
	body := &LoginRequestBody{
		Username: p.Username,
		Password: p.Password,
	}
	return body
}

// NewRegisterRequestBody builds the HTTP request body from the payload of the
// "register" endpoint of the "auth" service.
func NewRegisterRequestBody(p *auth.RegisterPayload) *RegisterRequestBody {
	body := &RegisterRequestBody{
		Username: p.Username,
		Password: p.Password,
	}
	return body
}

// NewLoginResultOK builds a "auth" service "login" endpoint result from a HTTP
// "OK" response.
func NewLoginResultOK(body *LoginResponseBody) *auth.LoginResult {
	v := &auth.LoginResult{
		Token: *body.Token,
	}

	return v
}

// NewLoginLoginFailed builds a auth service login endpoint login_failed error.
func NewLoginLoginFailed(body *LoginLoginFailedResponseBody) *auth.AuthFailed {
	v := &auth.AuthFailed{
		Message: *body.Message,
	}

	return v
}

// NewRegisterResultOK builds a "auth" service "register" endpoint result from
// a HTTP "OK" response.
func NewRegisterResultOK(body *RegisterResponseBody) *auth.RegisterResult {
	v := &auth.RegisterResult{
		Token: *body.Token,
	}

	return v
}

// NewRegisterRegisterFailed builds a auth service register endpoint
// register_failed error.
func NewRegisterRegisterFailed(body *RegisterRegisterFailedResponseBody) *auth.AuthFailed {
	v := &auth.AuthFailed{
		Message: *body.Message,
	}

	return v
}

// ValidateLoginResponseBody runs the validations defined on LoginResponseBody
func ValidateLoginResponseBody(body *LoginResponseBody) (err error) {
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	return
}

// ValidateRegisterResponseBody runs the validations defined on
// RegisterResponseBody
func ValidateRegisterResponseBody(body *RegisterResponseBody) (err error) {
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	return
}

// ValidateLoginLoginFailedResponseBody runs the validations defined on
// login_login_failed_response_body
func ValidateLoginLoginFailedResponseBody(body *LoginLoginFailedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateRegisterRegisterFailedResponseBody runs the validations defined on
// register_register_failed_response_body
func ValidateRegisterRegisterFailedResponseBody(body *RegisterRegisterFailedResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}
