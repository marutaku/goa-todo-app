// Code generated by goa v3.14.1, DO NOT EDIT.
//
// todo HTTP client types
//
// Command:
// $ goa gen backend/design

package client

import (
	todo "backend/gen/todo"

	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "todo" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// List of todos
	Todos TodoCollectionResponseBody `form:"todos,omitempty" json:"todos,omitempty" xml:"todos,omitempty"`
}

// ShowResponseBody is the type of the "todo" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// Todo to show
	Todo *TodoResponseBody `form:"todo,omitempty" json:"todo,omitempty" xml:"todo,omitempty"`
}

// TodoCollectionResponseBody is used to define fields on response body types.
type TodoCollectionResponseBody []*TodoResponseBody

// TodoResponseBody is used to define fields on response body types.
type TodoResponseBody struct {
	// Unique ID
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the todo
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the todo
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Whether or not the todo is done
	Done *bool `form:"done,omitempty" json:"done,omitempty" xml:"done,omitempty"`
	// When the todo was done in ISO format
	DoneAt *string `form:"doneAt,omitempty" json:"doneAt,omitempty" xml:"doneAt,omitempty"`
	// Who did the todo
	DoneBy *string `form:"doneBy,omitempty" json:"doneBy,omitempty" xml:"doneBy,omitempty"`
	// When the todo was created in ISO format
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Who created the todo
	CreatedBy *string `form:"createdBy,omitempty" json:"createdBy,omitempty" xml:"createdBy,omitempty"`
}

// NewListResultOK builds a "todo" service "list" endpoint result from a HTTP
// "OK" response.
func NewListResultOK(body *ListResponseBody) *todo.ListResult {
	v := &todo.ListResult{}
	if body.Todos != nil {
		v.Todos = make([]*todo.Todo, len(body.Todos))
		for i, val := range body.Todos {
			v.Todos[i] = unmarshalTodoResponseBodyToTodoTodo(val)
		}
	}

	return v
}

// NewShowResultOK builds a "todo" service "show" endpoint result from a HTTP
// "OK" response.
func NewShowResultOK(body *ShowResponseBody) *todo.ShowResult {
	v := &todo.ShowResult{}
	if body.Todo != nil {
		v.Todo = unmarshalTodoResponseBodyToTodoTodo(body.Todo)
	}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	if body.Todos != nil {
		if err2 := ValidateTodoCollectionResponseBody(body.Todos); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateShowResponseBody runs the validations defined on ShowResponseBody
func ValidateShowResponseBody(body *ShowResponseBody) (err error) {
	if body.Todo != nil {
		if err2 := ValidateTodoResponseBody(body.Todo); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTodoCollectionResponseBody runs the validations defined on
// TodoCollectionResponseBody
func ValidateTodoCollectionResponseBody(body TodoCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateTodoResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateTodoResponseBody runs the validations defined on TodoResponseBody
func ValidateTodoResponseBody(body *TodoResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Done == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("done", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdAt", "body"))
	}
	if body.CreatedBy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdBy", "body"))
	}
	return
}
