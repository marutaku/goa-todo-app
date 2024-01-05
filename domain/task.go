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
	DoneAt time.Time
	// Who did the todo
	DoneBy    string
	CreatedAt time.Time
	CreatedBy string
}

func NewTask(id TaskId, name string, description string, done bool, doneAtString string, oneBy string, createdAtString string, createdBy string) (*Task, error) {
	doneAt, err := time.Parse(time.RFC3339, doneAtString)
	if err != nil {
		return nil, err
	}
	createdAt, err := time.Parse(time.RFC3339, createdAtString)
	if err != nil {
		return nil, err
	}
	return &Task{
		ID:          id,
		Name:        name,
		Description: description,
		Done:        done,
		DoneAt:      doneAt,
		DoneBy:      oneBy,
		CreatedAt:   createdAt,
		CreatedBy:   createdBy,
	}, nil
}

func (t Task) IsSaved() bool {
	return t.ID != 0
}
