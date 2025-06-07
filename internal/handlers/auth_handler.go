package handlers

import (
	"service-store/internal/services"
	"service-store/internal/validators"
	"service-store/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	input, msg, errs := validators.ValidateBody[services.RegisterInput](c)

	if errs != nil {
		return utils.ErrorResponse(c, 403, msg, errs)
	} else if msg != "" {
		return utils.ErrorResponse(c, 403, msg, nil)
	}

	taken, err := services.IsEmailTaken(input.Email)

	if err != nil {
		return utils.ErrorResponse(c, 500, "Something went wrong.", nil)
	}

	if taken {
		ve := utils.NewValidationErrors()
		ve.Add("email", "Email already exists")
		return utils.ErrorResponse(c, 403, "Validation failed", ve.All())
	}

	if _, err := services.RegisterUser(*input); err != nil {
		return utils.ErrorResponse(c, 500, "Something went wrong.", nil)
	}

	return utils.CreatedResponse(c, "Register successfully", nil)

}

func Login(c *fiber.Ctx) error {

	input, msg, errs := validators.ValidateBody[services.LoginInput](c)
	if errs != nil {
		return utils.ErrorResponse(c, 422, msg, errs)
	}
	if msg != "" {
		return utils.ErrorResponse(c, 400, msg, nil)
	}

	user, err := services.IsExistUser(input.Email)
	if err != nil {
		return utils.ErrorResponse(c, 500, "Something went wrong.", nil)
	}
	if user == nil {
		return utils.ErrorResponse(c, 401, "Invalid credentials", nil)
	}

	if err := services.VerifyPassword(input.Password, user.Password); err != nil {
		return utils.ErrorResponse(c, 401, "Invalid credentials", nil)
	}

	token, err := services.GenerateJWTToken(user)
	if err != nil {
		return utils.ErrorResponse(c, 500, "Could not generate token.", nil)
	}

	return utils.SuccessResponse(c, "Login successfully", fiber.Map{
		"user":  user,
		"token": token,
	})
}

func ForgetPassword(c *fiber.Ctx) error {
	input, msg, errs := validators.ValidateBody[services.ForgetPasswordInput](c)

	if errs != nil {
		return utils.ErrorResponse(c, 422, msg, errs)
	}
	if msg != "" {
		return utils.ErrorResponse(c, 400, msg, nil)
	}

	if err := services.SendResetPasswordToken(input.Email); err != nil {
		return utils.ErrorResponse(c, 500, "Failed to send reset token", nil)
	}

	return utils.SuccessResponse(c, "Reset link sent successfully", nil)
}

func ResetPassword(c *fiber.Ctx) error {
	input, msg, errs := validators.ValidateBody[services.ResetPasswordInput](c)

	if errs != nil {
		return utils.ErrorResponse(c, 422, msg, errs)
	}
	if msg != "" {
		return utils.ErrorResponse(c, 400, msg, nil)
	}

	if err := services.ResetUserPassword(input.Token, input.NewPassword); err != nil {
		return utils.ErrorResponse(c, 400, err.Error(), nil)
	}

	return utils.SuccessResponse(c, "Password reset successfully", nil)
}
