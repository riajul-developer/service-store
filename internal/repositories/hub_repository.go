package repositories

import (
	"context"
	"service-store/config"
	"service-store/internal/models"
)

func CreateHub(hub *models.Hub) error {
	_, err := config.DB.NewInsert().Model(hub).Exec(context.Background())
	return err
}

func GetHubByID(id int) (*models.Hub, error) {
	hub := new(models.Hub)
	err := config.DB.NewSelect().
		Model(hub).
		Where("id = ?", id).
		Relation("HubIncharge").
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return hub, nil
}

func UpdateHub(hub *models.Hub) error {
	_, err := config.DB.NewUpdate().Model(hub).
		WherePK().
		Exec(context.Background())
	return err
}

func DeleteHub(id int) error {
	_, err := config.DB.NewDelete().
		Model((*models.Hub)(nil)).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}

func GetAllHubs(page, limit int) ([]models.Hub, int, error) {
	var hubs []models.Hub
	offset := (page - 1) * limit

	count, err := config.DB.NewSelect().
		Model(&hubs).
		Count(context.Background())
	if err != nil {
		return nil, 0, err
	}

	err = config.DB.NewSelect().
		Model(&hubs).
		Relation("HubIncharge").
		Limit(limit).
		Offset(offset).
		Scan(context.Background())
	if err != nil {
		return nil, 0, err
	}

	return hubs, count, nil
}
