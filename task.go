package taskapi

import (
	task "backend/gen/task"
	"backend/models"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// task service example implementation.
// The example methods log the requests and return zero values.
type tasksrvc struct {
	logger *log.Logger
	db     *gorm.DB
}

// NewTask returns the task service implementation.
func NewTask(logger *log.Logger) task.Service {
	config := models.NewPostgresConfig()
	db, err := models.NewPostgresDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	return &tasksrvc{logger, db}
}

// List all tasks
func (s *tasksrvc) List(ctx context.Context, p *task.ListPayload) (res *task.ListResult, err error) {
	res = &task.ListResult{}

	s.logger.Print("task.list")
	var tasks []*task.BackendStoredTask
	result := s.db.WithContext(ctx).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	res.Tasks = tasks
	return res, nil
}

// Show a task
func (s *tasksrvc) Show(ctx context.Context, p *task.ShowPayload) (res *task.ShowResult, err error) {
	res = &task.ShowResult{}
	s.logger.Print("task.show")

	var returnTask *task.BackendStoredTask
	result := s.db.WithContext(ctx).First(&returnTask, p.ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, task.NoMatch(fmt.Sprintf("no task found with id %d", p.ID))
		}
		return nil, result.Error
	}
	res.Task = returnTask
	return res, nil
}
