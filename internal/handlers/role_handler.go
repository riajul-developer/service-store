package handlers

import (
	"service-store/internal/services"
	"service-store/internal/validators"
	"service-store/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {

	input, msg, errs := validators.ValidateBody[services.CreateRoleInput](c)

	if errs != nil {
		return utils.ErrorResponse(c, 403, msg, errs)
	} else if msg != "" {
		return utils.ErrorResponse(c, 403, msg, nil)
	}

	taken, err := services.IsRoleNameTaken(input.Name)

	if err != nil {
		return utils.ErrorResponse(c, 500, "Something went wrong.", nil)
	}

	if taken {
		ve := utils.NewValidationErrors()
		ve.Add("name", "Name already exists")
		return utils.ErrorResponse(c, 403, "Validation failed", ve.All())
	}

	if role, err := services.CreateRole(*input); err != nil {
		return utils.ErrorResponse(c, 500, "Something went wrong.", nil)
	} else {
		return utils.CreatedResponse(c, "Register successfully", role)
	}

}

func AssignPermissions(c *fiber.Ctx) error {

	input, msg, errs := validators.ValidateBody[services.AssignPermissionInput](c)
	if errs != nil {
		return utils.ErrorResponse(c, 422, msg, errs)
	}
	if msg != "" {
		return utils.ErrorResponse(c, 400, msg, nil)
	}

	role, err := services.IsExistRole(input.RoleID)
	if err != nil {
		return utils.ErrorResponse(c, 500, "Something went wrong.", nil)
	}
	if role == nil {
		return utils.ErrorResponse(c, 401, "Invalid credentials", nil)
	}

	// // Verify password
	// if err := services.VerifyPassword(input.Password, user.Password); err != nil {
	// 	return utils.ErrorResponse(c, 401, "Invalid credentials", nil)
	// }

	// // Generate JWT token
	// token, err := services.GenerateJWTToken(user)
	// if err != nil {
	// 	return utils.ErrorResponse(c, 500, "Could not generate token.", nil)
	// }

	// // Return successful response
	return utils.SuccessResponse(c, "Login successfully", fiber.Map{
		"user":  "user",
		"token": "token",
	})
}
