package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()

	dsn := "host=localhost user=postgres password=sakib12345 dbname=mydb port=1234 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = d
	DB.AutoMigrate(&User{})

	e.GET("/users", GetAllUsers)
	e.POST("/users", CreateUser)
	e.DELETE("/users/:id", DeleteUser)

	e.Logger.Fatal(e.Start(":1111"))
}

func GetAllUsers(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := &User{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "User created successfully")
}

func DeleteUser(c echo.Context) error {
	var user User
	id := c.Param("id")

	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := DB.Where("id = ?", ID).Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}
