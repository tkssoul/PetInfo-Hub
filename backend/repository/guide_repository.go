package repository

import (
    "gorm.io/gorm"    
	"backend/models"
)

type GuideRepository struct {
    db *gorm.DB
}

// NewGuideRepository 创建一个新的 GuideRepository 实例
func NewGuideRepository(db *gorm.DB) *GuideRepository {
    return &GuideRepository{db: db}
}

// GetAllGuides 获取所有攻略
func (r *GuideRepository) GetAllGuides() ([]models.Guide, error) {
    var guides []models.Guide
    if err := r.db.Find(&guides).Error; err != nil {
        return nil, err
    }
    return guides, nil
}

// GetGuideByID 根据 ID 获取攻略
func (r *GuideRepository) GetGuideByID(guideID uint) (*models.Guide, error) {
    var guide models.Guide
    if err := r.db.First(&guide, guideID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil // 没有找到记录
        }
        return nil, err
    }
    return &guide, nil
}

// CreateGuide 创建新的攻略
func (r *GuideRepository) CreateGuide(guide models.Guide) (models.Guide, error) {
    if err := r.db.Create(&guide).Error; err != nil {
        return models.Guide{}, err
    }
    return guide, nil
}

// UpdateGuide 更新攻略信息
func (r *GuideRepository) UpdateGuide(guide models.Guide) error {
    if err := r.db.Save(&guide).Error; err != nil {
        return err
    }
    return nil
}

// DeleteGuide 删除攻略
func (r *GuideRepository) DeleteGuide(guideID uint) error {
    if err := r.db.Delete(&models.Guide{}, guideID).Error; err != nil {
        return err
    }
    return nil
}
