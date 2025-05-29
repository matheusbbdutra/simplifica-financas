package usecase

import (
	"errors"
	"simplificafinancas/internal/user/adapter/http/dto"
	"simplificafinancas/internal/user/domain"
	"simplificafinancas/pkg/utils"
)

type UpdateUserUseCase struct {
	Repo domain.UserRepository
}

func NewUpdateUserUseCase(repo domain.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		Repo: repo,
	}
}

func (uc *UpdateUserUseCase) Execute(email string, req *dto.UpdateUserRequest) (*domain.User, error) {
	user, err := uc.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}
    if req.Name != nil {
        user.Name = *req.Name
    }
    if req.Email != nil {
        user.Email = *req.Email
    }
    if req.Password != nil {
		password, err := utils.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
        user.Password = password
    }
    if err := uc.Repo.Update(user); err != nil {
        return nil, err
    }

	return user, nil
}