package handler

import (
	"simplificafinancas/internal/user/adapter/http/dto"
	"simplificafinancas/internal/user/application/usecase"
	"simplificafinancas/pkg/utils"

	"github.com/labstack/echo/v4"
)


type UserHandler struct {
	Validate *utils.Validate
	CreateUserUseCase *usecase.CreateUserUseCase
}

func NewUserHandler(Validate *utils.Validate, createUserUseCase *usecase.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		Validate: Validate,
		CreateUserUseCase: createUserUseCase,
	}
}


func (h *UserHandler) CreateUser(c echo.Context) error {

	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400,  map[string]string{"error": err.Error()})
	}

	if err := h.Validate.ValidateStruct(&req); err != nil {
		errors := h.Validate.TranslateError(err)
        return c.JSON(400, map[string]interface{}{"errors": errors})
	}

	err := h.CreateUserUseCase.Execute(&req)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, "User created successfully")
}


