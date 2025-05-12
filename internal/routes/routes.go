package routes

import (
	"service-store/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	role := app.Group("/roles")
	role.Post("/create", handlers.CreateRole)
	role.Post("/assign-permissions", handlers.AssignPermissions)

	// permission := app.Group("/permissions")
	// permission.Get("/", handlers.GetAllPermissions)
}
