package main

import (
	"github.com/gofiber/fiber/v2"
	// "service-store/config"
	"service-store/internal/routes"
	"service-store/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.ErrorResponse(c, 500, "Something went wrong", nil)
		},
	})
	// config.LoadEnv()
	// config.ConnectDB()

	routes.SetupRoutes(app)
	app.Listen(":3000")
}
