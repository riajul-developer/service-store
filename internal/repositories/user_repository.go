package repositories

import (
	"context"
	"service-store/config"
	"service-store/internal/models"
)

func CreateUser(user *models.User) error {
	_, err := config.DB.NewInsert().Model(user).Exec(context.Background())
	return err
}
