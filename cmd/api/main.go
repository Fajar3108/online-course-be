package main

import (
	"log"

	"github.com/Fajar3108/online-course-be/config"
	"github.com/Fajar3108/online-course-be/internal/router"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()

	app := router.SetupRoutes()

	app.Static("", "./public")

	err := app.Listen(":" + viper.GetString(config.AppPort))

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
