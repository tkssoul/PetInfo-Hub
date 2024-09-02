package repository

import (
    "errors"
    "gorm.io/gorm"
    "backend/models"
)

type SpotRepository struct {
    db *gorm.DB
}

// NewSpotRepository 创建一个新的 SpotRepository 实例
func NewSpotRepository(db *gorm.DB) *SpotRepository {
    return &SpotRepository{db: db}
}

// 创建宠物友好景点
func (r *SpotRepository) CreatePetFriendlySpot(spot *models.PetFriendlySpot) error {
    result := r.db.Create(spot)
    return result.Error
}

// 通过景点ID获取特定景点
func (r *SpotRepository) FindPetFriendlySpotByID(spotID uint) (*models.PetFriendlySpot, error) {
    var spot models.PetFriendlySpot
    result := r.db.First(&spot, spotID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("找不到该景点")
        }
        return nil, result.Error
    }
    return &spot, nil
}

// 获取所有宠物友好景点
func (r *SpotRepository) FindAllPetFriendlySpots() ([]models.PetFriendlySpot, error) {
    var spots []models.PetFriendlySpot
    result := r.db.Find(&spots)
    return spots, result.Error
}

// 更新宠物友好景点信息
func (r *SpotRepository) UpdatePetFriendlySpot(spot *models.PetFriendlySpot) error {
    result := r.db.Save(spot)
    return result.Error
}

// 删除特定宠物友好景点
func (r *SpotRepository) DeletePetFriendlySpot(spotID uint) error {
    result := r.db.Delete(&models.PetFriendlySpot{}, spotID)
    return result.Error
}
