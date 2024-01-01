// Code generated by goa v3.14.1, DO NOT EDIT.
//
// task HTTP client types
//
// Command:
// $ goa gen backend/design

package client

import (
	task "backend/gen/task"

	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "task" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// List of tasks
	Tasks BackendStoredTaskCollectionResponseBody `form:"tasks,omitempty" json:"tasks,omitempty" xml:"tasks,omitempty"`
}

// BackendStoredTaskCollectionResponseBody is used to define fields on response
// body types.
type BackendStoredTaskCollectionResponseBody []*BackendStoredTaskResponseBody

// BackendStoredTaskResponseBody is used to define fields on response body
// types.
type BackendStoredTaskResponseBody struct {
	// Unique ID
	ID *uint32 `gorm:"primaryKey"`
	// Name of the todo
	Name *string `gorm:"not null"`
	// Description of the todo
	Description *string `gorm:"not null;default ''"`
	// Whether or not the todo is done
	Done *bool `gorm:"not null;default false"`
	// When the todo was done in ISO format
	DoneAt *string `gorm:"default ''"`
	// Who did the todo
	DoneBy *string `gorm:"default ''"`
	// When the todo was created in ISO format
	CreatedAt *string `gorm:"autoCreateTime"`
	// Who created the todo
	CreatedBy *string `gorm:"not null"`
}

// NewListResultOK builds a "task" service "list" endpoint result from a HTTP
// "OK" response.
func NewListResultOK(body *ListResponseBody) *task.ListResult {
	v := &task.ListResult{}
	if body.Tasks != nil {
		v.Tasks = make([]*task.BackendStoredTask, len(body.Tasks))
		for i, val := range body.Tasks {
			v.Tasks[i] = unmarshalBackendStoredTaskResponseBodyToTaskBackendStoredTask(val)
		}
	}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	if body.Tasks != nil {
		if err2 := ValidateBackendStoredTaskCollectionResponseBody(body.Tasks); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateBackendStoredTaskCollectionResponseBody runs the validations defined
// on BackendStored-TaskCollectionResponseBody
func ValidateBackendStoredTaskCollectionResponseBody(body BackendStoredTaskCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateBackendStoredTaskResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateBackendStoredTaskResponseBody runs the validations defined on
// BackendStored-TaskResponseBody
func ValidateBackendStoredTaskResponseBody(body *BackendStoredTaskResponseBody) (err error) {
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
