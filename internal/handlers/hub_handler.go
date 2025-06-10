package handlers

import (
	"fmt"
	"service-store/internal/services"
	"service-store/internal/validators"
	"service-store/utils"

	"github.com/gofiber/fiber/v2"
)

// CreateHub handles hub creation
func CreateHub(c *fiber.Ctx) error {
	input, msg, errs := validators.ValidateBody[services.HubInput](c)
	if errs != nil {
		return utils.ErrorResponse(c, 403, msg, errs)
	} else if msg != "" {
		return utils.ErrorResponse(c, 403, msg, nil)
	}

	hub, err := services.CreateHub(*input)
	if err != nil {
		fmt.Println("CreateHub error:", err)
		return utils.ErrorResponse(c, 500, "Failed to create hub", hub)
	}

	return utils.CreatedResponse(c, "Hub created successfully", hub)
}

// GetAllHubs fetches all hubs
func GetAllHubs(c *fiber.Ctx) error {
	hubs, count, err := services.GetAllHubs(1, 10)
	if err != nil {
		return utils.ErrorResponse(c, 500, "Failed to fetch hubs", nil)
	}

	return utils.SuccessResponse(c, "Hubs fetched successfully", fiber.Map{
		"hubs":  hubs,
		"count": count,
	})
}

// GetHubByID fetches a single hub by ID
func GetHubByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, 400, "Invalid hub ID", nil)
	}

	hub, err := services.GetHubByID(id)
	if err != nil {
		return utils.ErrorResponse(c, 500, "Failed to fetch hub", nil)
	}
	if hub == nil {
		return utils.ErrorResponse(c, 404, "Hub not found", nil)
	}

	return utils.SuccessResponse(c, "Hub fetched successfully", hub)
}

// UpdateHub updates an existing hub
func UpdateHub(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, 400, "Invalid hub ID", nil)
	}

	input, msg, errs := validators.ValidateBody[services.HubInput](c)
	if errs != nil {
		return utils.ErrorResponse(c, 422, msg, errs)
	}
	if msg != "" {
		return utils.ErrorResponse(c, 400, msg, nil)
	}

	hub, err := services.UpdateHub(id, *input)
	if err != nil {
		return utils.ErrorResponse(c, 500, "Failed to update hub", nil)
	}
	if hub == nil {
		return utils.ErrorResponse(c, 404, "Hub not found", nil)
	}

	return utils.SuccessResponse(c, "Hub updated successfully", hub)
}

// DeleteHub deletes a hub by ID
func DeleteHub(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, 400, "Invalid hub ID", nil)
	}

	if err := services.DeleteHub(id); err != nil {
		return utils.ErrorResponse(c, 500, "Failed to delete hub", nil)
	}

	return utils.SuccessResponse(c, "Hub deleted successfully", nil)
}
