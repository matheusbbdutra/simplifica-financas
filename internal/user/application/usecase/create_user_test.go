package usecase_test

import (
    "testing"

    "simplificafinancas/internal/user/application/usecase"
    "simplificafinancas/internal/user/infrastructure/persistence"
    "simplificafinancas/internal/user/adapter/http/dto"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func setupTestUsecase(t *testing.T) *usecase.CreateUserUseCase {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }
    db.AutoMigrate(&persistence.UserModel{})
    repo := persistence.NewUserRepository(db)
    return usecase.NewCreateUserUseCase(repo)
}

func TestCreateUserUseCase_Execute(t *testing.T) {
    uc := setupTestUsecase(t)

    req := &dto.CreateUserRequest{
        Name:     "Usuário Teste",
        Email:    "teste@email.com",
        Password: "12345678",
    }

    err := uc.Execute(req)
    if err != nil {
		t.Errorf("esperado nil, obteve %v", err)
    }
}