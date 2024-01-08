package presenter

import (
	"backend/domain"
	taskService "backend/gen/task"
	"strconv"
	"time"
)

type TaskPresenter struct{}

func NewTaskPresenter() *TaskPresenter {
	return &TaskPresenter{}
}

func taskToOutput(task *domain.Task) *taskService.BackendStoredTask {
	doneAtString := ""
	if task.DoneAt != nil {
		doneAtString = task.DoneAt.Format(time.RFC3339)
	}
	createdAtString := task.CreatedAt.Format(time.RFC3339)
	storedTask := &taskService.BackendStoredTask{
		ID:          uint32(task.ID),
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
		DoneAt:      doneAtString,
		CreatedAt:   createdAtString,
		CreatedBy:   strconv.FormatUint(uint64(task.CreatedBy), 10),
	}
	if task.DoneBy != nil {
		storedTask.DoneBy = strconv.FormatUint(uint64(*task.DoneBy), 10)
	}
	return storedTask
}

func (p *TaskPresenter) ListOutput(tasks []*domain.Task) *taskService.ListResult {
	res := &taskService.ListResult{}
	var outputTasks []*taskService.BackendStoredTask
	for _, t := range tasks {
		outputTasks = append(outputTasks, taskToOutput(t))
	}
	res.Tasks = outputTasks
	return res
}

func (p *TaskPresenter) ShowOutput(task *domain.Task) *taskService.ShowResult {
	res := &taskService.ShowResult{}
	res.Task = taskToOutput(task)
	return res
}

func (p *TaskPresenter) CreateOutput(task *domain.Task) *taskService.CreateResult {
	res := &taskService.CreateResult{}
	res.Task = taskToOutput(task)
	return res
}

func (p *TaskPresenter) UpdateOutput(task *domain.Task) *taskService.UpdateResult {
	res := &taskService.UpdateResult{}
	res.Task = taskToOutput(task)
	return res
}

func (p *TaskPresenter) DoneOutput(task *domain.Task) *taskService.DoneResult {
	res := &taskService.DoneResult{}
	res.Task = taskToOutput(task)
	return res
}
