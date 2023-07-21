package user

import (
	"context"
	"go-learn/internal/delivery/http/dto"
	"go-learn/internal/domain"
)

func (uc UseCase) Create(ctx context.Context, dto dto.CreateUserDTO) (domain.User, error) {

	newUser := domain.NewUser(dto.Name)
	usrId, err := uc.storage.CreateUser(ctx, *newUser)

	if err != nil {
		return domain.User{}, err
	}

	newUser.ID = usrId
	return *newUser, nil
}
