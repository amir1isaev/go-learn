package storage

import (
	"context"
	"go-learn/internal/domain"

	"github.com/google/uuid"
)

type Storage interface {
	User
}

type User interface {
	CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error)
}
