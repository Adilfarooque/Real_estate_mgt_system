package repository

import "real_estate_mgt_system/internal/models"

type PropertyRepository interface {
	CreateProperty(property *models.Property)error
	GetPropertyByID(id uint) (*models.Property, error)
	GetPropertiesByID(userID uint) ([]*models.Property, error)
	UpdateProperty(property *models.Property) error
	DeleteProperty(id uint) error
}

