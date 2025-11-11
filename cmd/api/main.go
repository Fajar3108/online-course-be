package main

import (
	"log"

	"github.com/Fajar3108/online-course-be/config"
	"github.com/Fajar3108/online-course-be/internal/router"
)

func main() {
	app := router.SetupRoutes()

	app.Static("", "./public")

	err := app.Listen(":" + config.Config().App.Port)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
