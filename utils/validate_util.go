package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateBody[T any](c *fiber.Ctx) (*T, error) {
	var body T
	if err := c.BodyParser(&body); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validate.Struct(body); err != nil {
		var errs []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, map[string]string{
				"field":   e.Field(),
				"message": e.Tag(),
			})
		}

		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation failed",
			"data":    nil,
			"errors":  errs,
		})
	}

	return &body, nil
}
