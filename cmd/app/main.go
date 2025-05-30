// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"simplificafinancas/internal/app"
	_ "simplificafinancas/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("Arquivo .env não encontrado ou não pôde ser carregado")
    }

	e := echo.New()
	app := app.NewApp(e)
	app.InitDB()
	app.InitMiddleware()
	app.InitRoutes()

	e.Logger.Fatal(e.Start(":1323"))
}