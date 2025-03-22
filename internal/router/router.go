package router

import (
	"github.com/labstack/echo/v4"
	"goomers-api/internal/handler"
)

func InitUserRouter(e *echo.Echo) {
	e.GET("/users", handler.GetUsers)
	e.GET("/user/:userId", handler.GetUser)
	e.POST("/user", handler.AddUser)
}
