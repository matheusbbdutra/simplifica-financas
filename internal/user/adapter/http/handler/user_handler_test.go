package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"simplificafinancas/internal/user/adapter/http/dto"
	"simplificafinancas/internal/user/adapter/http/handler"
	"simplificafinancas/internal/user/application/usecase"
	"simplificafinancas/internal/user/infrastructure/persistence"
	"simplificafinancas/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func TestMain(m *testing.M) {
    _ = godotenv.Load(".env")
    os.Exit(m.Run())
}

func setupTestHandler(t *testing.T) *handler.UserHandler {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }
    db.AutoMigrate(&persistence.UserModel{})
    repo := persistence.NewUserRepository(db)
    createUser := usecase.NewCreateUserUseCase(repo)
    loginUser := usecase.NewLoginUserUseCase(repo)
    updateUser := usecase.NewUpdateUserUseCase(repo)
    validator := utils.NewValidator()
    return handler.NewUserHandler(validator, createUser, loginUser, updateUser)
}

func TestCreateUserHandler(t *testing.T) {
    e := echo.New()
    h := setupTestHandler(t)

    dto := dto.CreateUserRequest{
        Name:     "Teste",
        Email:    "teste@email.com",
        Password: "12345678",
    }   

    createTestUser(e, h, t, dto)
}

func TestLoginUserHandler(t *testing.T) {
    e := echo.New()
    h := setupTestHandler(t)

    createUser := dto.CreateUserRequest{
        Name:     "Teste",
        Email:    "teste@email.com",
        Password: "12345678",
    }  

    createTestUser(e, h, t, createUser)

    loginReqBody := dto.LoginUserRequest{
        Email:    "teste@email.com",
        Password: "12345678",
    }

    loginTestUser(e, h, t, loginReqBody)
}


func TestUpdateUserHandler(t *testing.T) {
    e := echo.New()
    h := setupTestHandler(t)

    createUser := dto.CreateUserRequest{
        Name:     "Teste",
        Email:    "teste@email.com",
        Password: "12345678",
    }
    createTestUser(e, h, t, createUser)
    
    loginReqBody := dto.LoginUserRequest{
        Email:    "teste@email.com",
        Password: "12345678",
    }
    token := loginTestUser(e, h, t, loginReqBody)
    
    name := "Teste Atualizado"
    email := "testeatualizado@email.com"
    password := "87654321"
    updateReqBody := dto.UpdateUserRequest{
        Name:     &name,
        Email:    &email,
        Password: &password,
    }
    body, _ := json.Marshal(updateReqBody)
    req := httptest.NewRequest(http.MethodPut, "/api/user/update", bytes.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    req.Header.Set("Authorization", "Bearer "+token)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    claims := jwt.MapClaims{
        "user_id": "1", 
        "email":   "teste@email.com",
    }
    tokenObj := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    c.Set("user", tokenObj)

    if err := h.UpdateUser(c); err != nil {
        t.Fatalf("erro ao atualizar usuário: %v", err)
    }
    if rec.Code != http.StatusOK {
        t.Errorf("esperado status 200 ao atualizar usuário, obteve %d", rec.Code)
    }
    repo := h.UpdateUserUseCase.Repo.(*persistence.UserRepositoryImpl)
    var user persistence.UserModel
    if err := repo.Db.Where("email = ?", email).First(&user).Error; err != nil {
        t.Fatalf("usuário não encontrado após update: %v", err)
    }
    if user.Name != name {
        t.Errorf("nome não atualizado, esperado %s, obteve %s", name, user.Name)
    }
    if user.Email != email {
        t.Errorf("email não atualizado, esperado %s, obteve %s", email, user.Email)
    }
    if user.Password == "" || !utils.CheckPasswordHash(password, user.Password) {
        t.Errorf("senha não atualizada, esperado %s, obteve %s", password, user.Password)
    }
} 

func createTestUser(e *echo.Echo, h *handler.UserHandler, t *testing.T, reqBody dto.CreateUserRequest) {
    body, _ := json.Marshal(reqBody)
    req := httptest.NewRequest(http.MethodPost, "/api/user/create", bytes.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    if err := h.CreateUser(c); err != nil {
        t.Fatalf("erro ao criar usuário: %v", err)
    }
    if rec.Code != http.StatusCreated {
        t.Errorf("esperado status 201 ao criar usuário, obteve %d", rec.Code)
    }
}

func loginTestUser(e *echo.Echo, h *handler.UserHandler, t *testing.T, reqBody dto.LoginUserRequest) string {
    body, _ := json.Marshal(reqBody)
    req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    if err := h.LoginUser(c); err != nil {
        t.Fatalf("erro ao fazer login: %v", err)
    }
    if rec.Code != http.StatusOK {
        t.Errorf("esperado status 200 ao logar, obteve %d", rec.Code)
    }
    var resp map[string]string
    if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
        t.Fatalf("erro ao decodificar resposta de login: %v", err)
    }
    token, ok := resp["token"]
    if !ok || token == "" {
        t.Fatalf("token JWT não retornado na resposta de login")
    }
    return token
}