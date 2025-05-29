package main

import (
	"simplificafinancas/internal/app"

	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()
	app := app.NewApp(e)
	app.InitDB()
	app.InitMiddleware()
	app.InitRoutes()

	e.Logger.Fatal(e.Start(":1323"))
}