package usecase

import (
    "simplificafinancas/internal/user/domain"
    "simplificafinancas/internal/user/adapter/http/dto"
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
    user := domain.NewUser(
        req.Name,
        req.Email,
        req.Password,
    )

    if err := uc.Repo.Create(user); err != nil {
        return err
    }

    return nil
}