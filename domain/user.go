package domain

import "time"

type UserId uint32

func (u UserId) UInt32() uint32 {
	return uint32(u)
}

type User struct {
	// Unique ID
	ID UserId
	// Name
	Name string
	// Hashed password of the user
	Password  string
	CreatedAt time.Time
}

func NewUser(id UserId, name string, password string, createdAt time.Time) (*User, error) {
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
