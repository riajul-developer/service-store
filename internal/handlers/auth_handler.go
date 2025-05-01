package handlers

import (
	"github.com/gofiber/fiber/v2"
	"service-store/internal/services"
	"service-store/internal/validators"
)

func Register(c *fiber.Ctx) error {
	var input services.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := validators.Validate.Struct(input)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.RegisterUser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}