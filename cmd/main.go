package main

import (
	"github.com/labstack/echo/v4"
	"goomers-api/internal/config"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})
	initializeDb()

	e.Logger.Fatal(e.Start(":8080"))
}

func initializeDb() {
	config.DatabaseInit()
	gorm := config.DB()
	doGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	gormErr := doGorm.Ping()
	if gormErr != nil {
		panic(gormErr)
	}
}
