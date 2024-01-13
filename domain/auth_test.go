package domain

import (
	"errors"
	"testing"
)

func TestAuthError(t *testing.T) {
	err := &AuthError{Err: errors.New("test")}
	if err.Error() != "test" {
		t.Error("Error is not test")
	}
}
