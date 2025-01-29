package repository

import (
	"real_estate_mgt_system/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByID(id models.User) (models.User, error)
	GetUserByEmail(email models.User) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User)error{
	return r.db.Create(&user).Error
}
