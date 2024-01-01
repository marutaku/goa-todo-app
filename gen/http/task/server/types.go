// Code generated by goa v3.14.1, DO NOT EDIT.
//
// task HTTP server types
//
// Command:
// $ goa gen backend/design

package server

import (
	task "backend/gen/task"
)

// ListResponseBody is the type of the "task" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// List of tasks
	Tasks BackendStoredTaskCollectionResponseBody `form:"tasks,omitempty" json:"tasks,omitempty" xml:"tasks,omitempty"`
}

// ShowResponseBody is the type of the "task" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// task to show
	Task *BackendStoredTaskResponseBody `form:"task,omitempty" json:"task,omitempty" xml:"task,omitempty"`
}

// BackendStoredTaskCollectionResponseBody is used to define fields on response
// body types.
type BackendStoredTaskCollectionResponseBody []*BackendStoredTaskResponseBody

// BackendStoredTaskResponseBody is used to define fields on response body
// types.
type BackendStoredTaskResponseBody struct {
	// Unique ID
	ID uint32 `gorm:"primaryKey"`
	// Name of the todo
	Name string `gorm:"not null"`
	// Description of the todo
	Description string `gorm:"not null;default ''"`
	// Whether or not the todo is done
	Done bool `gorm:"not null;default false"`
	// When the todo was done in ISO format
	DoneAt string `gorm:"default ''"`
	// Who did the todo
	DoneBy string `gorm:"default ''"`
	// When the todo was created in ISO format
	CreatedAt string `gorm:"autoCreateTime"`
	// Who created the todo
	CreatedBy string `gorm:"not null"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "task" service.
func NewListResponseBody(res *task.ListResult) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Tasks != nil {
		body.Tasks = make([]*BackendStoredTaskResponseBody, len(res.Tasks))
		for i, val := range res.Tasks {
			body.Tasks[i] = marshalTaskBackendStoredTaskToBackendStoredTaskResponseBody(val)
		}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "task" service.
func NewShowResponseBody(res *task.ShowResult) *ShowResponseBody {
	body := &ShowResponseBody{}
	if res.Task != nil {
		body.Task = marshalTaskBackendStoredTaskToBackendStoredTaskResponseBody(res.Task)
	}
	return body
}

// NewListPayload builds a task service list endpoint payload.
func NewListPayload(limit uint32, offset uint32) *task.ListPayload {
	v := &task.ListPayload{}
	v.Limit = limit
	v.Offset = offset

	return v
}

// NewShowPayload builds a task service show endpoint payload.
func NewShowPayload(id uint32) *task.ShowPayload {
	v := &task.ShowPayload{}
	v.ID = id

	return v
}
