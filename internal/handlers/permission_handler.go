package handlers

// import (
// 	"service-store/internal/services"

// 	"github.com/gofiber/fiber/v2"
// )

// // GetAllPermissions handles GET /permissions
// func GetAllPermissions(c *fiber.Ctx) error {
// 	permissions, err := services.GetAllPermissions()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to retrieve permissions",
// 			"error":   err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message":     "Permissions retrieved successfully",
// 		"permissions": permissions,
// 	})
// }
