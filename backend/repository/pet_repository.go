package repository

import (
    "errors"
    "backend/models"
    "gorm.io/gorm"
)

type PetRepository struct {
    db *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
    return &PetRepository{db: db}
}

// 创建宠物
func (r *PetRepository) CreatePet(pet *models.Pets) error {
    result := r.db.Create(pet)
    return result.Error
}

// 通过ID查找宠物
func (r *PetRepository) FindPetByID(petID uint) (*models.Pets, error) {
    var pet models.Pets
    result := r.db.First(&pet, petID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("宠物不存在")
        }
        return nil, result.Error
    }
    return &pet, nil
}

// 更新宠物信息
func (r *PetRepository) UpdatePet(pet *models.Pets) error {
    result := r.db.Save(pet)
    return result.Error
}

// 删除宠物
func (r *PetRepository) DeletePet(petID uint) error {
    result := r.db.Delete(&models.Pets{}, petID)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("宠物不存在")
    }
    return nil
}

// 获取特定用户的宠物列表
func (r *PetRepository) GetPetsByUserID(userID uint) ([]models.Pets, error) {
    var pets []models.Pets
    result := r.db.Where("user_id = ?", userID).Find(&pets)
    if result.Error != nil {
        return nil, result.Error
    }
    return pets, nil
}
