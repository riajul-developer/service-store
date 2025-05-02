package services

import (
	// "golang.org/x/crypto/bcrypt"
	// "service-store/internal/models"
	// "service-store/internal/repositories"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// func RegisterUser(input RegisterInput) error {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
// 	user := &models.User{
// 		Name:     input.Name,
// 		Email:    input.Email,
// 		Password: string(hashedPassword),
// 	}
// 	return repositories.CreateUser(user)
// }