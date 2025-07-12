package routes

import (
	"github.com/labstack/echo/v4"
	"main.go/internal/controllers"
)

func UserRoutes(e *echo.Echo) {
	u := e.Group("/app")
	u.POST("/users", controllers.CreateUser)
	u.GET("/users", controllers.GetUser)
	u.GET("/users/:id", controllers.GetUserByID)
	u.PUT("/users/:id", controllers.UpdateUser)
	u.DELETE("/users/:id", controllers.DeleteUser)
}
