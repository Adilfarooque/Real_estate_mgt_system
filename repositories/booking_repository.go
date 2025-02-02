package repository

import (
	"real_estate_mgt_system/internal/models"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking)
	GetBookingByID(id uint) (*models.Booking, error)
	GetBookingByUserID(userID uint) ([]*models.Booking, error)
	UpdateBooking(booking *models.Booking)
	DeleteBooking(id uint)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingByID(id uint) (*models.Booking,error){
	var booking models.Booking
	err := r.db.First(&booking, id).Error
}