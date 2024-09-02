package repository

import (
    "errors"
    "gorm.io/gorm"
    "backend/models"
)

type PetBoardingDetailRepository struct {
    db *gorm.DB
}

// NewPetBoardingDetailRepository 创建一个新的 PetBoardingDetailRepository 实例
func NewPetBoardingDetailRepository(db *gorm.DB) *PetBoardingDetailRepository {
    return &PetBoardingDetailRepository{db: db}
}

// 创建寄养信息
func (r *PetBoardingDetailRepository) CreatePetBoardingDetail(detail *models.PetBoardingDetail) error {
    result := r.db.Create(detail)
    return result.Error
}

// 通过寄养信息ID获取特定寄养信息
func (r *PetBoardingDetailRepository) FindPetBoardingDetailByID(boardingID int) (*models.PetBoardingDetail, error) {
    var detail models.PetBoardingDetail
    result := r.db.First(&detail, boardingID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("找不到该寄养信息")
        }
        return nil, result.Error
    }
    return &detail, nil
}

// 获取特定寄养人的所有寄养信息
func (r *PetBoardingDetailRepository) FindPetBoardingDetailsBySitterID(sitterID int) ([]models.PetBoardingDetail, error) {
    var details []models.PetBoardingDetail
    result := r.db.Where("sitter_id = ?", sitterID).Find(&details)
    return details, result.Error
}

// 更新寄养信息
func (r *PetBoardingDetailRepository) UpdatePetBoardingDetail(detail *models.PetBoardingDetail) error {
    result := r.db.Save(detail)
    return result.Error
}

// 删除特定寄养信息
func (r *PetBoardingDetailRepository) DeletePetBoardingDetail(boardingID int) error {
    result := r.db.Delete(&models.PetBoardingDetail{}, boardingID)
    return result.Error
}
