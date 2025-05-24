package handler_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "simplificafinancas/internal/user/adapter/http/dto"
    "simplificafinancas/internal/user/adapter/http/handler"
    "simplificafinancas/internal/user/application/usecase"
    "simplificafinancas/internal/user/infrastructure/persistence"
    "simplificafinancas/pkg/utils"

    "github.com/labstack/echo/v4"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func setupTestHandler(t *testing.T) *handler.UserHandler {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }
    db.AutoMigrate(&persistence.UserModel{})
    repo := persistence.NewUserRepository(db)
    usecase := usecase.NewCreateUserUseCase(repo)
    validator := utils.NewValidator()
    return handler.NewUserHandler(validator, usecase)
}

func TestCreateUserHandler(t *testing.T) {
    e := echo.New()
    h := setupTestHandler(t)

    reqBody := dto.CreateUserRequest{
        Name:     "Teste",
        Email:    "teste@email.com",
        Password: "12345678",
    }
    body, _ := json.Marshal(reqBody)
    req := httptest.NewRequest(http.MethodPost, "/api/user/create", bytes.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if err := h.CreateUser(c); err != nil {
        t.Fatalf("handler retornou erro: %v", err)
    }


    if rec.Code != http.StatusCreated {
        t.Errorf("esperado status 201, obteve %d", rec.Code)
    }
}