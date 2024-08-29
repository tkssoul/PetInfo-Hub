package services

import (
    "backend/models"
)

func GetNearbyHotels() ([]models.Hotel, error) {
    var hotels []models.Hotel
    result := models.DB.Find(&hotels)
    return hotels, result.Error
}
