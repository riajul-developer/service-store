package main

import (
	"github.com/gofiber/fiber/v2"
	"service-store/config"
	"service-store/internal/routes"
)

func main() {
	app := fiber.New()
	config.LoadEnv()
	config.ConnectDB()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}