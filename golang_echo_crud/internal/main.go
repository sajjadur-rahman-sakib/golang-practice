package main

import (
	"github.com/labstack/echo/v4"
	"main.go/command/config"
	"main.go/command/routes"
)

func main() {
	e := echo.New()
	config.ConnectDB()
	routes.UserRoutes(e)
	e.Logger.Fatal(e.Start(":1111"))
}
