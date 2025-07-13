package main

import (
	"golang_login_signup/internal/config"
	"golang_login_signup/internal/models"
	"golang_login_signup/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	e := echo.New()
	routes.AuthRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
