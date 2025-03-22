package handler

import (
	"github.com/labstack/echo/v4"
	"goomers-api/internal/model"
	"goomers-api/internal/service"
	"net/http"
)

var usrService *service.UserService

func InitUserHandler(us *service.UserService) {
	usrService = us
}
func GetUsers(c echo.Context) error {

	users, err := usrService.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	userId := c.Param("userId")
	user, err := usrService.GetUser(userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	err := usrService.AddUser(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}
