package taskapi

import (
	task "backend/gen/task"
	"backend/infrastructure/database"
	"context"
	"fmt"
	"log"
	"time"

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
	config := database.NewPostgresConfig()
	db, err := database.NewPostgresDatabase(config)
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
	criteria := map[string]interface{}{}
	if p.CreatedBy != "" {
		criteria["created_by"] = p.CreatedBy
	}
	if p.Name != "" {
		criteria["name"] = p.Name
	}
	result := s.db.WithContext(ctx).Where(criteria).Order("created_at DESC").Limit(int(p.Limit)).Offset(int(p.Offset)).Find(&tasks)
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

func (s *tasksrvc) Create(ctx context.Context, p *task.CreatePayload) (res *task.CreateResult, err error) {
	res = &task.CreateResult{}
	s.logger.Print("task.create")
	newTask := &task.BackendStoredTask{
		Name:        p.Name,
		Description: *p.Description,
		Done:        false,
		DoneAt:      "",
		DoneBy:      "",
		CreatedBy:   p.CreatedBy,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	result := s.db.WithContext(ctx).Create(&newTask)
	if result.Error != nil {
		return nil, result.Error
	}
	res.Task = newTask
	fmt.Println(newTask)
	return res, nil
}

func (s *tasksrvc) Update(ctx context.Context, p *task.UpdatePayload) (res *task.UpdateResult, err error) {
	res = &task.UpdateResult{}
	s.logger.Print("task.update")
	var returnTask *task.BackendStoredTask
	result := s.db.WithContext(ctx).First(&returnTask, p.ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, task.NoMatch(fmt.Sprintf("no task found with id %d", p.ID))
		}
		return nil, result.Error
	}
	if p.Name != nil {
		returnTask.Name = *p.Name
	}
	if p.Description != nil {
		returnTask.Description = *p.Description
	}
	result = s.db.WithContext(ctx).Save(&returnTask)
	if result.Error != nil {
		return nil, result.Error
	}
	res.Task = returnTask
	return res, nil
}

func (s *tasksrvc) Done(ctx context.Context, p *task.DonePayload) (res *task.DoneResult, err error) {
	res = &task.DoneResult{}
	s.logger.Print("task.done")
	var returnTask *task.BackendStoredTask
	result := s.db.WithContext(ctx).First(&returnTask, p.ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, task.NoMatch(fmt.Sprintf("no task found with id %d", p.ID))
		}
		return nil, result.Error
	}
	returnTask.Done = true
	returnTask.DoneAt = time.Now().Format(time.RFC3339)
	returnTask.DoneBy = *p.DoneBy
	result = s.db.WithContext(ctx).Save(&returnTask)
	if result.Error != nil {
		return nil, result.Error
	}
	res.Task = returnTask
	return res, nil
}
