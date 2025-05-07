package utils

import (
	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

// Success Responses
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func NoContentResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNoContent).JSON(APIResponse{
		Status:  "success",
		Message: message,
	})
}

// Generic Error Response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, errs interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Status:  "error",
		Message: message,
		Errors:  errs,
	})
}

// Error Shortcut Wrappers
// func BadRequest(c *fiber.Ctx, message string, errs interface{}) error {
// 	return ErrorResponse(c, fiber.StatusBadRequest, message, errs)
// }

// func NotFound(c *fiber.Ctx, message string) error {
// 	return ErrorResponse(c, fiber.StatusNotFound, message, nil)
// }

// func ServerError(c *fiber.Ctx, message string) error {
// 	return ErrorResponse(c, fiber.StatusInternalServerError, message, nil)
// }

// func Unauthorized(c *fiber.Ctx, message string) error {
// 	return ErrorResponse(c, fiber.StatusUnauthorized, message, nil)
// }

// func Forbidden(c *fiber.Ctx, message string) error {
// 	return ErrorResponse(c, fiber.StatusForbidden, message, nil)
// }

// func Conflict(c *fiber.Ctx, message string, errs interface{}) error {
// 	return ErrorResponse(c, fiber.StatusConflict, message, errs)
// }
