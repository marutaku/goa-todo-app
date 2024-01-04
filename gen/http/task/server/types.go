// Code generated by goa v3.14.1, DO NOT EDIT.
//
// task HTTP server types
//
// Command:
// $ goa gen backend/design

package server

import (
	task "backend/gen/task"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "task" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// ID of task to create
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the task
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Who created the task
	CreatedBy *string `form:"created_by,omitempty" json:"created_by,omitempty" xml:"created_by,omitempty"`
}

// UpdateRequestBody is the type of the "task" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Name of the task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the task
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// DoneRequestBody is the type of the "task" service "done" endpoint HTTP
// request body.
type DoneRequestBody struct {
	// Who did the task
	DoneBy *string `form:"done_by,omitempty" json:"done_by,omitempty" xml:"done_by,omitempty"`
}

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

// CreateResponseBody is the type of the "task" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// Created task
	Task *BackendStoredTaskResponseBody `form:"task,omitempty" json:"task,omitempty" xml:"task,omitempty"`
}

// UpdateResponseBody is the type of the "task" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// Updated task
	Task *BackendStoredTaskResponseBody `form:"task,omitempty" json:"task,omitempty" xml:"task,omitempty"`
}

// DoneResponseBody is the type of the "task" service "done" endpoint HTTP
// response body.
type DoneResponseBody struct {
	// Finished task
	Task *BackendStoredTaskResponseBody `form:"task,omitempty" json:"task,omitempty" xml:"task,omitempty"`
}

// BackendStoredTaskCollectionResponseBody is used to define fields on response
// body types.
type BackendStoredTaskCollectionResponseBody []*BackendStoredTaskResponseBody

// BackendStoredTaskResponseBody is used to define fields on response body
// types.
type BackendStoredTaskResponseBody struct {
	// Unique ID
	ID uint32 `gorm:"primaryKey" json:"id"`
	// Name of the todo
	Name string `gorm:"not null" json:"name"`
	// Description of the todo
	Description string `gorm:"not null;default ''" json:"description"`
	// Whether or not the todo is done
	Done bool `gorm:"not null;default false" json:"done"`
	// When the todo was done in ISO format
	DoneAt string `gorm:"default ''" json:"doneAt"`
	// Who did the todo
	DoneBy string `gorm:"default ''" json:"doneBy"`
	// When the todo was created in ISO format
	CreatedAt string `gorm:"not null" json:"createdAt"`
	// Who created the todo
	CreatedBy string `gorm:"not null" json:"createdBy"`
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

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "task" service.
func NewCreateResponseBody(res *task.CreateResult) *CreateResponseBody {
	body := &CreateResponseBody{}
	if res.Task != nil {
		body.Task = marshalTaskBackendStoredTaskToBackendStoredTaskResponseBody(res.Task)
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "task" service.
func NewUpdateResponseBody(res *task.UpdateResult) *UpdateResponseBody {
	body := &UpdateResponseBody{}
	if res.Task != nil {
		body.Task = marshalTaskBackendStoredTaskToBackendStoredTaskResponseBody(res.Task)
	}
	return body
}

// NewDoneResponseBody builds the HTTP response body from the result of the
// "done" endpoint of the "task" service.
func NewDoneResponseBody(res *task.DoneResult) *DoneResponseBody {
	body := &DoneResponseBody{}
	if res.Task != nil {
		body.Task = marshalTaskBackendStoredTaskToBackendStoredTaskResponseBody(res.Task)
	}
	return body
}

// NewListPayload builds a task service list endpoint payload.
func NewListPayload(limit uint32, offset uint32, createdBy string, name string) *task.ListPayload {
	v := &task.ListPayload{}
	v.Limit = limit
	v.Offset = offset
	v.CreatedBy = createdBy
	v.Name = name

	return v
}

// NewShowPayload builds a task service show endpoint payload.
func NewShowPayload(id uint32) *task.ShowPayload {
	v := &task.ShowPayload{}
	v.ID = id

	return v
}

// NewCreatePayload builds a task service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody) *task.CreatePayload {
	v := &task.CreatePayload{
		ID:          body.ID,
		Name:        *body.Name,
		Description: body.Description,
		CreatedBy:   *body.CreatedBy,
	}

	return v
}

// NewUpdatePayload builds a task service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id uint32) *task.UpdatePayload {
	v := &task.UpdatePayload{
		Name:        body.Name,
		Description: body.Description,
	}
	v.ID = id

	return v
}

// NewDonePayload builds a task service done endpoint payload.
func NewDonePayload(body *DoneRequestBody, id uint32) *task.DonePayload {
	v := &task.DonePayload{
		DoneBy: body.DoneBy,
	}
	v.ID = &id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.CreatedBy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_by", "body"))
	}
	return
}
