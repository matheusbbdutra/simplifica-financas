package app

import (
	mid "simplificafinancas/internal/common/adapters/http/middleware"
	"simplificafinancas/internal/user/adapter/http/handler"
	"simplificafinancas/internal/user/application/usecase"
	"simplificafinancas/internal/user/infrastructure/persistence"
	"simplificafinancas/pkg/utils"

	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Echo *echo.Echo
	db   *gorm.DB
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

func (a *App) InitMiddleware() {
	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())
	a.Echo.Use(mid.JWTAuth())

}

func (a *App) InitRoutes() {
	a.newUserModule()
	a.Echo.GET("/public/swagger/*", echoSwagger.WrapHandler)

}

func (a *App) newUserModule() {
	userRepo := persistence.NewUserRepository(a.db)
	createUser := usecase.NewCreateUserUseCase(userRepo)
	loginUser := usecase.NewLoginUserUseCase(userRepo)
	updateUser := usecase.NewUpdateUserUseCase(userRepo)
	validate := utils.NewValidator()
	handler := handler.NewUserHandler(validate, createUser, loginUser, updateUser)
	userGroup := a.Echo.Group("/api/user")

	userGroup.POST("/register", handler.CreateUser)
	userGroup.POST("/login", handler.LoginUser)
	userGroup.POST("/update", handler.UpdateUser)
}
