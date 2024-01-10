package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"context"
	"time"
)

type (
	TaskUseCase interface {
		List(ctx context.Context, criteria TaskCriteria) ([]*domain.Task, error)
		Show(ctx context.Context, id uint32) (*domain.Task, error)
		Create(ctx context.Context, params TaskCreateParams) (*domain.Task, error)
		Update(ctx context.Context, id uint32, params TaskUpdateParams) (*domain.Task, error)
		Delete(ctx context.Context, id uint32) error
		Done(ctx context.Context, id uint32) (*domain.Task, error)
	}
	TaskCriteria struct {
		Name        *string
		Description *string
		Done        *bool
		CreatedBy   domain.UserId
	}
	TaskCreateParams struct {
		Name        string
		Description string
	}
	TaskUpdateParams struct {
		Name        *string
		Description *string
	}
	taskInteractor struct {
		taskRepo repository.TaskRepositoryInterface
		userRepo repository.UserRepositoryInterface
	}
)

func NewTaskInteractor(taskRepo repository.TaskRepositoryInterface, userRepo repository.UserRepositoryInterface) *taskInteractor {
	return &taskInteractor{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (t *taskInteractor) List(ctx context.Context, criteria TaskCriteria) ([]*domain.Task, error) {
	return t.taskRepo.FindAll(ctx, criteria.Name, criteria.Done, criteria.CreatedBy)
}

func (t *taskInteractor) Show(ctx context.Context, id uint32) (*domain.Task, error) {
	return t.taskRepo.FindOne(ctx, domain.TaskId(id))
}

func (t *taskInteractor) Create(ctx context.Context, params TaskCreateParams) (*domain.Task, error) {
	task, err := domain.NewTask(
		domain.TaskId(0),
		params.Name,
		params.Description,
		false,
		nil,
		nil,
		time.Now(),
		ctx.Value("userId").(domain.UserId),
	)
	if err != nil {
		return nil, err
	}
	return t.taskRepo.Create(ctx, task)
}

func (t *taskInteractor) Update(ctx context.Context, id uint32, params TaskUpdateParams) (*domain.Task, error) {
	taskId := domain.TaskId(id)
	task, err := t.taskRepo.FindOne(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if params.Name != nil {
		task.Name = *params.Name
	}
	if params.Description != nil {
		task.Description = *params.Description
	}
	return t.taskRepo.Update(ctx, task)
}

func (t *taskInteractor) Delete(ctx context.Context, id uint32) error {
	taskId := domain.TaskId(id)
	task, err := t.taskRepo.FindOne(ctx, taskId)
	if err != nil {
		return err
	}
	return t.taskRepo.Delete(ctx, task.ID)
}

func (t *taskInteractor) Done(ctx context.Context, id uint32) (*domain.Task, error) {
	taskId := domain.TaskId(id)
	task, err := t.taskRepo.FindOne(ctx, taskId)
	if err != nil {
		return nil, err
	}
	task.Done = true
	doneAt := time.Now()
	task.DoneAt = &doneAt
	task.DoneBy = ctx.Value("userId").(*domain.UserId)
	return t.taskRepo.Update(ctx, task)
}
