package routes

import (
	"golang_login_signup/internal/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	a := e.Group("/auth")
	a.POST("/signup", controllers.Signup)
	a.POST("/login", controllers.Login)
}
