package usecase

import (
	"errors"
	"fmt"
	"simplificafinancas/internal/user/adapter/http/dto"
	"simplificafinancas/internal/user/domain"
	"simplificafinancas/pkg/utils"
)

type LoginUserUseCase struct {
	Repo domain.UserRepository
}

func NewLoginUserUseCase(repo domain.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{
		Repo: repo,
	}
}

func (uc *LoginUserUseCase) Execute(dto dto.LoginUserRequest) (*string, error) {
	user, err := uc.Repo.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user with email %s not found", dto.Email)
	}

	if !user.CheckPassword(dto.Password) {
		return nil, errors.New("invalid password") 
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &token, nil
}