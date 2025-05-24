package app

import (
	"simplificafinancas/internal/user/adapter/http/handler"
	"simplificafinancas/internal/user/application/usecase"
	"simplificafinancas/internal/user/infrastructure/persistence"
	"simplificafinancas/pkg/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Echo *echo.Echo
	db *gorm.DB
	// User
}

func NewApp(e *echo.Echo) *App {
	return &App{Echo: e}
}

func (a *App) InitDB() {
    db, err := gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
	db.AutoMigrate(&persistence.UserModel{})
    a.db = db
}

func (a *App) InitRoutes() {
	a.newUserModule()

}

func (a *App) newUserModule() {
	userRepo := persistence.NewUserRepository(a.db)
	createUser := usecase.NewCreateUserUseCase(userRepo)
	validate := utils.NewValidator()
	handler := handler.NewUserHandler(validate, createUser)
	userGroup := a.Echo.Group("/api/user")
	
	userGroup.POST("/create", handler.CreateUser)
}