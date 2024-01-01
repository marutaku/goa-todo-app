// Code generated by goa v3.14.1, DO NOT EDIT.
//
// HTTP request path constructors for the task service.
//
// Command:
// $ goa gen backend/design

package client

import (
	"fmt"
)

// ListTaskPath returns the URL path to the task service list HTTP endpoint.
func ListTaskPath() string {
	return "/tasks"
}

// ShowTaskPath returns the URL path to the task service show HTTP endpoint.
func ShowTaskPath(id uint32) string {
	return fmt.Sprintf("/tasks/%v", id)
}

// CreateTaskPath returns the URL path to the task service create HTTP endpoint.
func CreateTaskPath() string {
	return "/tasks"
}
