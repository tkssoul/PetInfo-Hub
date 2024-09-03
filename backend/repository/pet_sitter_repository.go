package repository

import (
    "errors"
    "gorm.io/gorm"
    "backend/models"
)

type PetSitterRepository struct {
    db *gorm.DB
}

// NewPetSitterRepository 创建一个新的 PetSitterRepository 实例
func NewPetSitterRepository(db *gorm.DB) *PetSitterRepository {
    return &PetSitterRepository{db: db}
}

// 创建寄养人
func (r *PetSitterRepository) CreatePetSitter(sitter *models.PetSitter) error {
    result := r.db.Create(sitter)
    return result.Error
}

// 通过寄养人ID获取特定寄养人
func (r *PetSitterRepository) FindPetSitterByID(sitterID uint) (*models.PetSitter, error) {
    var sitter models.PetSitter
    result := r.db.First(&sitter, sitterID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("找不到该寄养人")
        }
        return nil, result.Error
    }
    return &sitter, nil
}

// 获取所有寄养人
func (r *PetSitterRepository) FindAllPetSitters() ([]models.PetSitter, error) {
    var sitters []models.PetSitter
    result := r.db.Find(&sitters)
    return sitters, result.Error
}

// 更新寄养人信息
func (r *PetSitterRepository) UpdatePetSitter(sitter *models.PetSitter) error {
    result := r.db.Save(sitter)
    return result.Error
}

// 删除特定寄养人
func (r *PetSitterRepository) DeletePetSitter(sitterID uint) error {
    result := r.db.Delete(&models.PetSitter{}, sitterID)
    return result.Error
}
