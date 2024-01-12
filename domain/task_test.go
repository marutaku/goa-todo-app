package domain

import (
	"testing"
	"time"
)

func TestTaskNewTask(t *testing.T) {
	createdTime := time.Now()
	task, err := NewTask(1, "test", "test", false, nil, nil, createdTime, 1)
	if err != nil {
		t.Error(err)
	}
	if task.ID != 1 {
		t.Error("ID is not 1")
	}
	if task.Name != "test" {
		t.Error("Name is not test")
	}
	if task.Description != "test" {
		t.Error("Description is not test")
	}
	if task.Done != false {
		t.Error("Done is not false")
	}
	if task.DoneAt != nil {
		t.Error("DoneAt is not nil")
	}
	if task.DoneBy != nil {
		t.Error("DoneBy is not nil")
	}
	if task.CreatedAt != createdTime {
		t.Error("CreatedAt is not nil")
	}
	if task.CreatedBy != 1 {
		t.Error("CreatedBy is not 1")
	}
}

func TestIsSaved(t *testing.T) {
	task, err := NewTask(1, "test", "test", false, nil, nil, time.Now(), 1)
	if err != nil {
		t.Error(err)
	}
	if task.IsSaved() != true {
		t.Error("IsSaved is not true")
	}
	task.ID = 0
	if task.IsSaved() != false {
		t.Error("IsSaved is not false")
	}
}
