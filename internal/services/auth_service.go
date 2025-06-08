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

type ForgetPasswordInput struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordInput struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

var secret = func() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}
	return secret
}()

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func SendResetPasswordToken(email string) error {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	siteUrl := os.Getenv("SITE_URL")

	resetLink := siteUrl + "/reset-password?token=" + signedToken

	return utils.SendEmail(user.Email, "Password Reset", "reset_password.html", map[string]interface{}{
		"ResetLink":   resetLink,
		"AppName":     "Service Store",
		"UserName":    user.Name,
		"CurrentYear": time.Now().Year(),
	})

}

func ResetUserPassword(tokenStr string, newPassword string) error {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		return jwt.ErrInvalidKey
	}

	userID := int64(claims["user_id"].(float64))
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	user.Password = string(hashedPassword)

	return repositories.UpdateUserPassword(user)
}
