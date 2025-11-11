package router

import (
	"github.com/Fajar3108/online-course-be/config"
	errorhandler "github.com/Fajar3108/online-course-be/pkg/error-handler"
	"github.com/Fajar3108/online-course-be/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func SetupRoutes() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorhandler.GlobalErrorHandler,
	})

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.Config().CookieSecretKey,
	}))

	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to API")
	})

	AuthRouter(api.Group("/auth"))

	api.Use(middleware.JWTMiddleware())
	CategoryRouter(api)

	return app
}
