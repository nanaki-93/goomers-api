package config

import (
	"github.com/labstack/echo/v4"
	"goomers-api/internal/handler"
	"goomers-api/internal/repository"
	"goomers-api/internal/router"
	"goomers-api/internal/service"
)

func InitServer() *echo.Echo {

	e := echo.New()

	dbConn := NewDatabaseConnection()

	userRepo := repository.NewUserRepository(dbConn)

	userService := service.NewUserService(userRepo)

	handler.InitUserHandler(userService)
	router.InitUserRouter(e)

	return e
}
