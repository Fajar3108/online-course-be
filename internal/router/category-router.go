package router

import (
	"github.com/Fajar3108/online-course-be/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func CategoryRouter(r fiber.Router) {
	categoryController := controller.NewCategoryController()

	r.Get("/categories", categoryController.Index)
	r.Get("/categories/:slug", categoryController.Show)
	r.Post("/categories", categoryController.Store)
	r.Patch("/categories/:slug", categoryController.Update)
	r.Delete("/categories/:slug", categoryController.Destroy)
}
