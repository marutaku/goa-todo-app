package domain

import "time"

type TaskId uint32

func (t TaskId) UInt32() uint32 {
	return uint32(t)
}

type Task struct {
	// Unique ID
	ID TaskId
	// Name of the todo
	Name string
	// Description of the todo
	Description string
	// Whether or not the todo is done
	Done bool
	// When the todo was done in ISO format
	DoneAt *time.Time
	// Who did the todo
	DoneBy    string
	CreatedAt time.Time
	CreatedBy string
}

func NewTask(id TaskId, name string, description string, done bool, doneAt *time.Time, doneBy string, createdAt time.Time, createdBy string) (*Task, error) {
	return &Task{
		ID:          id,
		Name:        name,
		Description: description,
		Done:        done,
		DoneAt:      doneAt,
		DoneBy:      doneBy,
		CreatedAt:   createdAt,
		CreatedBy:   createdBy,
	}, nil
}

func (t Task) IsSaved() bool {
	return t.ID != 0
}
