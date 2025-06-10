package services

import (
	"service-store/internal/models"
	"service-store/internal/repositories"
)

type HubInput struct {
	Name          string  `json:"name" validate:"required,min=3"`
	Description   *string `json:"description" validate:"omitempty,min=3"`
	Address       string  `json:"address" validate:"required"`
	Latitude      float64 `json:"latitude" validate:"required"`
	Longitude     float64 `json:"longitude" validate:"required"`
	HubInchargeID int     `json:"hub_incharge_id" validate:"required"`
	IsActive      bool    `json:"is_active"`
}

func CreateHub(input HubInput) (*models.Hub, error) {
	hub := &models.Hub{
		Name:          input.Name,
		Description:   input.Description,
		Address:       input.Address,
		Latitude:      input.Latitude,
		Longitude:     input.Longitude,
		HubInchargeID: input.HubInchargeID,
		IsActive:      input.IsActive,
	}

	if err := repositories.CreateHub(hub); err != nil {
		return nil, err
	}

	return hub, nil
}

func GetHubByID(id int) (*models.Hub, error) {
	return repositories.GetHubByID(id)
}

func UpdateHub(id int, input HubInput) (*models.Hub, error) {
	hub, err := repositories.GetHubByID(id)
	if err != nil {
		return nil, err
	}

	hub.Name = input.Name
	hub.Description = input.Description
	hub.Address = input.Address
	hub.Latitude = input.Latitude
	hub.Longitude = input.Longitude
	hub.HubInchargeID = input.HubInchargeID
	hub.IsActive = input.IsActive

	if err := repositories.UpdateHub(hub); err != nil {
		return nil, err
	}

	return hub, nil
}

func DeleteHub(id int) error {
	return repositories.DeleteHub(id)
}

func GetAllHubs(page, limit int) ([]models.Hub, int, error) {
	return repositories.GetAllHubs(page, limit)
}
