package taskapi

import (
	task "backend/gen/task"
	"context"
	"log"
)

// task service example implementation.
// The example methods log the requests and return zero values.
type tasksrvc struct {
	logger *log.Logger
}

// NewTask returns the task service implementation.
func NewTask(logger *log.Logger) task.Service {
	return &tasksrvc{logger}
}

// List all tasks
func (s *tasksrvc) List(ctx context.Context, p *task.ListPayload) (res *task.ListResult, err error) {
	res = &task.ListResult{}
	s.logger.Print("task.list")
	return
}

// Show a task
func (s *tasksrvc) Show(ctx context.Context, p *task.ShowPayload) (res *task.ShowResult, err error) {
	res = &task.ShowResult{}
	s.logger.Print("task.show")
	return
}
