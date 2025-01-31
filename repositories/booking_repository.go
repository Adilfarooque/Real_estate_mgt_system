package repository

import "real_estate_mgt_system/internal/models"

type BookingRepository interface {
	CreateBooking(booking *models.Booking)
	GetBookingByID(id uint)(*models.Booking,error)
	GetBookingByUserID(userID uint)([]*models.Booking,error)
	UpdateBooking(booking *models.Booking)
	DeleteBooking(id uint)
}

