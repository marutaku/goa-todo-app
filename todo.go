package todoapi

import (
	todo "backend/gen/todo"
	"context"
	"log"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger *log.Logger
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger) todo.Service {
	return &todosrvc{logger}
}

// List all todos
func (s *todosrvc) List(ctx context.Context, p *todo.ListPayload) (res *todo.ListResult, err error) {
	res = &todo.ListResult{
		Todos: []*todo.Task{
			{
				ID:          1,
				Name:        "task1",
				Description: "description1",
				Done:        true,
				DoneAt:      "2023-12-015T11:13:30.143822",
				DoneBy:      "doneBy",
				CreatedAt:   "2023-12-01T11:13:30.143822",
				CreatedBy:   "tmpUser",
			},
			{
				ID:          2,
				Name:        "task2",
				Description: "description2",
				Done:        false,
				DoneAt:      "",
				DoneBy:      "",
				CreatedAt:   "2023-12-01T11:13:30.143822",
				CreatedBy:   "tmpUser",
			},
		},
	}
	return
}

func (s *todosrvc) Show(ctx context.Context, p *todo.ShowPayload) (res *todo.ShowResult, err error) {
	res = &todo.ShowResult{
		Todo: &todo.Task{
			ID:          1,
			Name:        "task1",
			Description: "description1",
			Done:        true,
			DoneAt:      "2023-12-015T11:13:30.143822",
			DoneBy:      "doneBy",
			CreatedAt:   "2023-12-01T11:13:30.143822",
			CreatedBy:   "tmpUser",
		},
	}
	return
}
