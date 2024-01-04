// Code generated by goa v3.14.1, DO NOT EDIT.
//
// task endpoints
//
// Command:
// $ goa gen backend/design

package task

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "task" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Show   goa.Endpoint
	Create goa.Endpoint
	Update goa.Endpoint
	Done   goa.Endpoint
}

// NewEndpoints wraps the methods of the "task" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:   NewListEndpoint(s),
		Show:   NewShowEndpoint(s),
		Create: NewCreateEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Done:   NewDoneEndpoint(s),
	}
}

// Use applies the given middleware to all the "task" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Create = m(e.Create)
	e.Update = m(e.Update)
	e.Done = m(e.Done)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "task".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
		return s.List(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "task".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ShowPayload)
		return s.Show(ctx, p)
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "task".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
		return s.Create(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "task".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
		return s.Update(ctx, p)
	}
}

// NewDoneEndpoint returns an endpoint function that calls the method "done" of
// service "task".
func NewDoneEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DonePayload)
		return s.Done(ctx, p)
	}
}
