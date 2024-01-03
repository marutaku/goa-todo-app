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

// CreateRequestBody is the type of the "task" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// ID of task to create
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the task
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the task
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Who created the task
	CreatedBy string `form:"created_by" json:"created_by" xml:"created_by"`
}

// UpdateRequestBody is the type of the "task" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Name of the task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the task
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
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
	CreatedAt *string `gorm:"not null"`
	// Who created the todo
	CreatedBy *string `gorm:"not null"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "task" service.
func NewCreateRequestBody(p *task.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedBy:   p.CreatedBy,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "task" service.
func NewUpdateRequestBody(p *task.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name:        p.Name,
		Description: p.Description,
	}
	return body
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

// NewShowResultOK builds a "task" service "show" endpoint result from a HTTP
// "OK" response.
func NewShowResultOK(body *ShowResponseBody) *task.ShowResult {
	v := &task.ShowResult{}
	if body.Task != nil {
		v.Task = unmarshalBackendStoredTaskResponseBodyToTaskBackendStoredTask(body.Task)
	}

	return v
}

// NewShowNoMatch builds a task service show endpoint no_match error.
func NewShowNoMatch(body string) task.NoMatch {
	v := task.NoMatch(body)

	return v
}

// NewCreateResultCreated builds a "task" service "create" endpoint result from
// a HTTP "Created" response.
func NewCreateResultCreated(body *CreateResponseBody) *task.CreateResult {
	v := &task.CreateResult{}
	if body.Task != nil {
		v.Task = unmarshalBackendStoredTaskResponseBodyToTaskBackendStoredTask(body.Task)
	}

	return v
}

// NewUpdateResultOK builds a "task" service "update" endpoint result from a
// HTTP "OK" response.
func NewUpdateResultOK(body *UpdateResponseBody) *task.UpdateResult {
	v := &task.UpdateResult{}
	if body.Task != nil {
		v.Task = unmarshalBackendStoredTaskResponseBodyToTaskBackendStoredTask(body.Task)
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

// ValidateShowResponseBody runs the validations defined on ShowResponseBody
func ValidateShowResponseBody(body *ShowResponseBody) (err error) {
	if body.Task != nil {
		if err2 := ValidateBackendStoredTaskResponseBody(body.Task); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.Task != nil {
		if err2 := ValidateBackendStoredTaskResponseBody(body.Task); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateResponseBody runs the validations defined on UpdateResponseBody
func ValidateUpdateResponseBody(body *UpdateResponseBody) (err error) {
	if body.Task != nil {
		if err2 := ValidateBackendStoredTaskResponseBody(body.Task); err2 != nil {
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
