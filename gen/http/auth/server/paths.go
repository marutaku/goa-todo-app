// Code generated by goa v3.14.1, DO NOT EDIT.
//
// HTTP request path constructors for the auth service.
//
// Command:
// $ goa gen backend/design

package server

// LoginAuthPath returns the URL path to the auth service login HTTP endpoint.
func LoginAuthPath() string {
	return "/login"
}

// RegisterAuthPath returns the URL path to the auth service register HTTP endpoint.
func RegisterAuthPath() string {
	return "/register"
}

// LogoutAuthPath returns the URL path to the auth service logout HTTP endpoint.
func LogoutAuthPath() string {
	return "/logout"
}
