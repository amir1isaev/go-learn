package usecase

import (
	"context"
	"go-learn/internal/delivery/http/dto"
	"go-learn/internal/domain"
)

type User interface {
	Create(ctx context.Context, dto dto.CreateUserDTO) (domain.User, error)
}
