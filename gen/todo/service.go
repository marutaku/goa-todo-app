// Code generated by goa v3.14.1, DO NOT EDIT.
//
// todo service
//
// Command:
// $ goa gen backend/design

package todo

import (
	"context"
)

// The todo service manages todo lists
type Service interface {
	// List all todos
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// Show a todo
	Show(context.Context, *ShowPayload) (res *ShowResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "todo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"list", "show"}

// ListPayload is the payload type of the todo service list method.
type ListPayload struct {
	// Maximum number of todos to return
	Limit uint32
	// Offset into the list of todos to start at
	Offset uint32
}

// ListResult is the result type of the todo service list method.
type ListResult struct {
	// List of todos
	Todos TodoCollection
}

// ShowPayload is the payload type of the todo service show method.
type ShowPayload struct {
	// ID of todo to show
	ID uint32
}

// ShowResult is the result type of the todo service show method.
type ShowResult struct {
	// Todo to show
	Todo *Todo
}

// Todo describes a todo item
type Todo struct {
	// Unique ID
	ID uint32
	// Name of the todo
	Name string
	// Description of the todo
	Description string
	// Whether or not the todo is done
	Done bool
	// When the todo was done in ISO format
	DoneAt string
	// Who did the todo
	DoneBy string
	// When the todo was created in ISO format
	CreatedAt string
	// Who created the todo
	CreatedBy string
}

type TodoCollection []*Todo
