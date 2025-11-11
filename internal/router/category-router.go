package router

import (
	"github.com/Fajar3108/online-course-be/database"
	"github.com/Fajar3108/online-course-be/internal/controller"
	"github.com/Fajar3108/online-course-be/internal/service"
	"github.com/gofiber/fiber/v2"
)

func CategoryRouter(r fiber.Router) {
	categoryService := service.NewCategoryService(database.DB())
	categoryController := controller.NewCategoryController(categoryService)

	r.Get("/categories", categoryController.Index)
	r.Get("/categories/:slug", categoryController.Show)
	r.Post("/categories", categoryController.Store)
	r.Patch("/categories/:slug", categoryController.Update)
	r.Delete("/categories/:slug", categoryController.Destroy)
}
