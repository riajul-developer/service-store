package services

import (
	"os"
	"service-store/internal/models"
	"service-store/internal/repositories"
	"service-store/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterUser handles the user registration
func RegisterUser(input RegisterInput) (*models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	if err := repositories.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

// IsEmailTaken checks if the email already exists
func IsEmailTaken(email string) (bool, error) {
	user, err := repositories.GetUserByEmail(email)

	if err != nil {
		if !utils.IsNotFoundError(err) {
			return false, err
		}
		return false, nil
	}

	return user != nil, nil
}

// IsExistUser checks if a user exists by email and returns the user if found
func IsExistUser(email string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(email)

	if err != nil {
		if !utils.IsNotFoundError(err) {
			return nil, err
		}
		return nil, nil
	}

	return user, nil
}

// VerifyPassword compares the provided password with the stored password
func VerifyPassword(providedPassword, storedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
}

// GenerateJWTToken creates a signed JWT for the given user
func GenerateJWTToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get secret key from environment variable or default
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}

	// Sign the token
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
