package repository

import (
    "backend/models"
    "gorm.io/gorm"
)

type HotelRepository struct {
    db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *HotelRepository {
    return &HotelRepository{db: db}
}

func (r *HotelRepository) FindAllHotels() ([]models.Hotel, error) {
    var hotels []models.Hotel
    result := r.db.Find(&hotels)
    return hotels, result.Error
}
