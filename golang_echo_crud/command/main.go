package main

import (
	"github.com/labstack/echo/v4"
	"main.go/internal/config"
	"main.go/internal/routes"
)

func main() {
	e := echo.New()
	config.ConnectDB()
	routes.UserRoutes(e)
	e.Logger.Fatal(e.Start(":1111"))
}
