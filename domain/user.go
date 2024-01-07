package domain

import "time"

type User struct {
	// Unique ID
	ID uint32
	// Name
	Name string
	// Hashed password of the user
	Password  string
	CreatedAt time.Time
}

func NewUser(id uint32, name string, password string, createdAt time.Time) (*User, error) {
	return &User{
		ID:        id,
		Name:      name,
		Password:  password,
		CreatedAt: createdAt,
	}, nil
}

func (u *User) IsSaved() bool {
	return u.ID != 0
}
