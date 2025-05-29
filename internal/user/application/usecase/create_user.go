package usecase

import (
	"simplificafinancas/internal/user/adapter/http/dto"
	"simplificafinancas/internal/user/domain"
	"simplificafinancas/pkg/utils"
)

type CreateUserUseCase struct {
    Repo domain.UserRepository
}

func NewCreateUserUseCase(repo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		Repo: repo,
	}
}

func (uc *CreateUserUseCase) Execute(req *dto.CreateUserRequest) (error) {
    password, err := utils.HashPassword(req.Password)
    
    if err != nil {
        return err
    }

    user := domain.NewUser(
        req.Name,
        req.Email,
        password,
    )

    if err := uc.Repo.Create(user); err != nil {
        return err
    }

    return nil
}