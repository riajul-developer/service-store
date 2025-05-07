package handlers

import (
	"service-store/internal/services"
	"service-store/internal/validators"
	"service-store/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// var input services.RegisterInput
	// if err := c.BodyParser(&input); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	// }

	// err := validators.Validate.Struct(input)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	// }

	input, msg, errs := validators.ValidateBody[services.RegisterInput](c)

	if errs != nil {
		return utils.ErrorResponse(c, 403, msg, errs)
	} else if msg != "" {
		return utils.ErrorResponse(c, 403, msg, nil)
	}

	if err := services.RegisterUser(*input); err != nil {
		return utils.ErrorResponse(c, 403, "msg", nil)
	}

	return utils.SuccessResponse(c, "Register successfully", nil)

}
