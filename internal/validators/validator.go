package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return "Invalid email address"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}

func ValidateBody[T any](c *fiber.Ctx) (*T, string, []map[string]string) {
	var body T

	if err := c.BodyParser(&body); err != nil {
		return nil, "Invalid request body", nil
	}

	if err := validate.Struct(body); err != nil {
		var errs []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, map[string]string{
				"field":   e.Field(),
				"message": getErrorMessage(e),
			})
		}

		return nil, "Validation failed", errs
	}

	return &body, "", nil
}
