package user

import (
	"go-learn/internal/usecase/adapter/storage"
)

type UseCase struct {
	storage storage.User
}

func NewUseCase(storage storage.User) *UseCase {
	return &UseCase{
		storage: storage,
	}
}
