package main

import (
	"log"

	"github.com/Fajar3108/online-course-be/database"
	"github.com/Fajar3108/online-course-be/pkg/model"
)

func main() {
	err := database.DB().AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.UserSession{},
	)

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}
