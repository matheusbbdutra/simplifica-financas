package handler

import (
	"simplificafinancas/internal/user/adapter/http/dto"
	"simplificafinancas/internal/user/application/usecase"
	"simplificafinancas/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


type UserHandler struct {
	Validate *utils.Validate
	CreateUserUseCase *usecase.CreateUserUseCase
	LoginUserUseCase *usecase.LoginUserUseCase
	UpdateUserUseCase *usecase.UpdateUserUseCase
}

func NewUserHandler(Validate *utils.Validate, createUserUseCase *usecase.CreateUserUseCase, loginUserUseCase *usecase.LoginUserUseCase, updateUserUseCase *usecase.UpdateUserUseCase) *UserHandler {
	return &UserHandler{
		Validate: Validate,
		CreateUserUseCase: createUserUseCase,
		LoginUserUseCase: loginUserUseCase,
		UpdateUserUseCase: updateUserUseCase,
	}
}

// CreateUser godoc
// @Summary Cria um novo usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user  body  dto.CreateUserRequest  true  "Dados do usuário"
// @Success 201 {object} string "User created successfully"
// @Failure 400 {object} map[string]any
// @Router /api/user/register [post]
func (h *UserHandler) CreateUser(c echo.Context) error {

	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400,  map[string]string{"error": err.Error()})
	}

	if err := h.Validate.ValidateStruct(&req); err != nil {
		errors := h.Validate.TranslateError(err)
        return c.JSON(400, map[string]any{"errors": errors})
	}

	err := h.CreateUserUseCase.Execute(&req)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, "User created successfully")
}

// LoginUser godoc
// @Summary Realiza login do usuário
// @Description Autentica o usuário e retorna um token JWT
// @Tags user
// @Accept  json
// @Produce  json
// @Param   credentials  body  dto.LoginUserRequest  true  "Credenciais do usuário"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]string
// @Router /api/user/login [post]
func (h *UserHandler) LoginUser(c echo.Context) error {
	var req dto.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.Validate.ValidateStruct(&req); err != nil {
		errors := h.Validate.TranslateError(err)
		return c.JSON(400, map[string]interface{}{"errors": errors})
	}

	token, err := h.LoginUserUseCase.Execute(req)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"token": *token})
}

// UpdateUser godoc
// @Summary Atualiza dados do usuário
// @Description Atualiza as informações do usuário autenticado
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user  body  dto.UpdateUserRequest  true  "Dados para atualização"
// @Success 200 {string} string "user updated successfully"
// @Failure 400 {object} map[string]any
// @Failure 401 {object} map[string]any
// @Router /api/user/update [post]
// @Security ApiKeyAuth
func (h *UserHandler) UpdateUser(c echo.Context) error {

	user := c.Get("user")
    if user == nil {
        return c.JSON(401, map[string]string{"message": "Unauthorized"})
    }

    userToken := user.(*jwt.Token)
    claims := userToken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.Validate.ValidateStruct(&req); err != nil {
		errors := h.Validate.TranslateError(err)
		return c.JSON(400, map[string]any{"errors": errors})
	}

	user, err := h.UpdateUserUseCase.Execute(email, &req)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, "user updated successfully")
}