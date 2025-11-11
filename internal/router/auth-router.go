package router

import (
	"github.com/Fajar3108/online-course-be/database"
	"github.com/Fajar3108/online-course-be/internal/controller"
	"github.com/Fajar3108/online-course-be/internal/service"
	"github.com/Fajar3108/online-course-be/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(r fiber.Router) {
	authService := service.NewAuthService(database.DB())
	authController := controller.NewAuthController(authService)

	r.Post("/login", authController.Login)
	r.Post("/register", authController.Register)
	r.Put("/refresh-token", authController.RefreshToken)

	r.Use(middleware.JWTMiddleware())
	r.Delete("/logout", authController.Logout)
}
