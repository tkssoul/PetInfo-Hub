package repository

import (
    "errors"
    "backend/models"
    "gorm.io/gorm"
)

type SpotRepository struct {
    db *gorm.DB
}

func NewSpotRepository(db *gorm.DB) *SpotRepository {
    return &SpotRepository{db: db}
}

// CreatePetFriendlySpot 创建宠物友好景点
func (r *SpotRepository) CreatePetFriendlySpot(spot *models.PetFriendlySpots) error {
    result := r.db.Create(spot)
    return result.Error
}

// FindSpotByID 通过景点ID查找景点
func (r *SpotRepository) FindSpotByID(spotID int64) (*models.PetFriendlySpots, error) {
    var spot models.PetFriendlySpots
    result := r.db.First(&spot, spotID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("景点不存在")
        }
        return nil, result.Error
    }
    return &spot, nil
}

// FindAllSpots 获取所有宠物友好景点
func (r *SpotRepository) FindAllSpots() ([]models.PetFriendlySpots, error) {
    var spots []models.PetFriendlySpots
    result := r.db.Find(&spots)
    if result.Error != nil {
        return nil, result.Error
    }
    return spots, nil
}

// UpdatePetFriendlySpot 更新宠物友好景点信息
func (r *SpotRepository) UpdatePetFriendlySpot(spot *models.PetFriendlySpots) error {
    result := r.db.Save(spot)
    return result.Error
}

// DeletePetFriendlySpot 删除宠物友好景点
func (r *SpotRepository) DeletePetFriendlySpot(spotID int64) error {
    result := r.db.Delete(&models.PetFriendlySpots{}, spotID)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("景点不存在")
    }
    return nil
}
