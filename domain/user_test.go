package domain

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	currentTime := time.Now()
	user, err := NewUser(1, "test", "test", currentTime)
	if err != nil {
		t.Error(err)
	}
	if user.ID != 1 {
		t.Error("ID is not 1")
	}
	if user.Name != "test" {
		t.Error("Name is not test")
	}
	if user.Password != "test" {
		t.Error("Password is not test")
	}
	if user.CreatedAt != currentTime {
		t.Error("CreatedAt is not nil")
	}
}

func TestUserIsSaved(t *testing.T) {
	user, err := NewUser(1, "test", "test", time.Now())
	if err != nil {
		t.Error(err)
	}
	if user.IsSaved() != true {
		t.Error("IsSaved is not true")
	}
	user.ID = 0
	if user.IsSaved() != false {
		t.Error("IsSaved is not false")
	}
}
