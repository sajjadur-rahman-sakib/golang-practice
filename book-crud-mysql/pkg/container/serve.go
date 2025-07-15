package container

import (
	"book-crud/pkg/config"
	"book-crud/pkg/connection"
	"book-crud/pkg/controllers"
	"book-crud/pkg/repositories"
	"book-crud/pkg/routes"
	"book-crud/pkg/service"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	config.SetConfig()

	db := connection.GetDB()

	bookRepo := repositories.BookDBInstance(db)

	bookService := service.BookServiceInstance(bookRepo)

	controllers.SetBookService(bookService)

	routes.BookRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
