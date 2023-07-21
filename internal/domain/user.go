package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewUser(
	name string,
) *User {
	return &User{
		Name: name,
	}
}
