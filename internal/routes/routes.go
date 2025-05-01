package routes

import (
	"github.com/gofiber/fiber/v2"
	"service-store/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
}
