package repository

import (
    "backend/models"
    "gorm.io/gorm"
)

type PetFriendlySpotRepository struct {
    db *gorm.DB
}

func NewPetFriendlySpotRepository(db *gorm.DB) *PetFriendlySpotRepository {
    return &PetFriendlySpotRepository{db: db}
}

// 获取附近的宠物友好景点
func (r *PetFriendlySpotRepository) FindAllHotels() ([]models.PetFriendlySpot, error) {
    var hotels []models.PetFriendlySpot
    result := r.db.Find(&hotels)
    return hotels, result.Error
}

// 通过ID查找宠物友好景点
func (r *PetFriendlySpotRepository) FindHotelByID(postID uint) (*models.PetFriendlySpot, error) {
    var Spot *models.PetFriendlySpot
    result := r.db.First(&Spot, postID)
    return Spot, result.Error
}



