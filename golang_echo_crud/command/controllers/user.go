package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/command/config"
	"main.go/command/models"
)

func CreateUser(e echo.Context) error {
	user := new(models.User)

	if err := e.Bind(user); err != nil {
		return err
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, user)
}

func GetUser(e echo.Context) error {
	var users []models.User

	config.DB.Find(&users)
	return e.JSON(http.StatusOK, users)
}

func GetUserByID(e echo.Context) error {
	id := e.Param("id")
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return e.JSON(http.StatusOK, user)
}

func UpdateUser(e echo.Context) error {
	id := e.Param("id")
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{
			"error": "User not found",
		})
	}
	if err := e.Bind(&user); err != nil {
		return err
	}
	config.DB.Save(&user)
	return e.JSON(http.StatusOK, user)
}

func DeleteUser(e echo.Context) error {
	id := e.Param("id")
	if err := config.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return e.JSON(http.StatusOK, echo.Map{"message": "User deleted"})
}
