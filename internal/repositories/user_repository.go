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

func GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := config.DB.NewSelect().Model(user).Where("email = ?", email).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(id int64) (*models.User, error) {
	user := new(models.User)
	err := config.DB.NewSelect().Model(user).Where("id = ?", id).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserPassword(user *models.User) error {
	_, err := config.DB.NewUpdate().
		Model(user).
		Column("password").
		Where("id = ?", user.ID).
		Exec(context.Background())
	return err
}
