package controller

import (
	"backend/adapter/presenter"
	"backend/adapter/repository"
	"backend/domain"
	taskService "backend/gen/task"
	"backend/infrastructure/database"
	"backend/service"
	"backend/usecase"
	"context"
	"fmt"
	"log"

	"goa.design/goa/v3/security"
	"gorm.io/gorm"
)

// task service example implementation.
// The example methods log the requests and return zero values.
type taskController struct {
	logger      *log.Logger
	usecase     usecase.TaskUseCase
	authService *service.JWTAuthService
	presenter   *presenter.TaskPresenter
}

// NewTask returns the task service implementation.
func NewTaskController(logger *log.Logger) *taskController {
	config := database.NewPostgresConfig()
	db, err := database.NewPostgresDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	taskRepository := repository.NewTaskRepository(db, logger)
	userRepository := repository.NewUserRepository(db, logger)
	return &taskController{
		logger:      logger,
		usecase:     usecase.NewTaskInteractor(taskRepository, userRepository),
		authService: service.NewJwTAuthService(userRepository),
		presenter:   presenter.NewTaskPresenter(),
	}
}

func (c *taskController) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	userId, err := c.authService.VerifyToken(ctx, token)
	if err != nil {
		return nil, &taskService.AuthFailed{Message: err.Error()}
	}
	ctx = context.WithValue(ctx, "userId", userId)
	return ctx, nil
}

// List all tasks
func (c *taskController) List(ctx context.Context, p *taskService.ListPayload) (res *taskService.ListResult, err error) {
	c.logger.Print("task.list")
	criteria := usecase.TaskCriteria{
		CreatedBy: domain.UserId(ctx.Value("userId").(uint32)),
	}
	if p.Name != "" {
		criteria.Name = &p.Name
	}
	tasks, err := c.usecase.List(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return c.presenter.ListOutput(tasks), nil
}

// Show a task
func (c *taskController) Show(ctx context.Context, p *taskService.ShowPayload) (res *taskService.ShowResult, err error) {
	task, err := c.usecase.Show(ctx, p.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, taskService.NoMatch(fmt.Sprintf("no task found with id %d", p.ID))
		}
		return nil, err
	}
	return c.presenter.ShowOutput(task), nil
}

func (c *taskController) Create(ctx context.Context, p *taskService.CreatePayload) (res *taskService.CreateResult, err error) {
	c.logger.Print("task.create")
	task, err := c.usecase.Create(ctx, usecase.TaskCreateParams{
		Name:        p.Name,
		Description: *p.Description,
	})
	if err != nil {
		return nil, err
	}
	return c.presenter.CreateOutput(task), nil
}

func (c *taskController) Update(ctx context.Context, p *taskService.UpdatePayload) (res *taskService.UpdateResult, err error) {
	c.logger.Print("task.update")
	task, err := c.usecase.Update(ctx, p.ID, usecase.TaskUpdateParams{
		Name:        p.Name,
		Description: p.Description,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, taskService.NoMatch(fmt.Sprintf("no task found with id %d", p.ID))
		}
		return nil, err
	}
	return c.presenter.UpdateOutput(task), nil
}

func (c *taskController) Done(ctx context.Context, p *taskService.DonePayload) (res *taskService.DoneResult, err error) {
	c.logger.Print("task.done")
	task, err := c.usecase.Done(ctx, p.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, taskService.NoMatch(fmt.Sprintf("no task found with id %d", p.ID))
		}
		return nil, err
	}
	return c.presenter.DoneOutput(task), nil
}

func (c *taskController) Delete(ctx context.Context, p *taskService.DeletePayload) (err error) {
	c.logger.Print("task.delete")
	err = c.usecase.Delete(ctx, p.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return taskService.NoMatch(fmt.Sprintf("no task found with id %d", p.ID))
		}
		return err
	}
	return nil
}
