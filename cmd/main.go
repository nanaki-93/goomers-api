package main

import (
	"goomers-api/internal/config"
)

func main() {

	e := config.InitServer()

	e.Logger.Fatal(e.Start(":8080"))
}
