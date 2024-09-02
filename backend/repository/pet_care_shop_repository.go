package repository

import (
    "errors"
    "gorm.io/gorm"
    "backend/models"
)

type PetCareShopRepository struct {
    db *gorm.DB
}

// NewPetCareShopRepository 创建一个新的 PetCareShopRepository 实例
func NewPetCareShopRepository(db *gorm.DB) *PetCareShopRepository {
    return &PetCareShopRepository{db: db}
}

// 创建宠物服务店铺
func (r *PetCareShopRepository) CreatePetCareShop(shop *models.PetCareShop) error {
    result := r.db.Create(shop)
    return result.Error
}

// 通过店铺ID获取特定店铺
func (r *PetCareShopRepository) FindPetCareShopByID(shopID int) (*models.PetCareShop, error) {
    var shop models.PetCareShop
    result := r.db.First(&shop, shopID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("找不到该店铺")
        }
        return nil, result.Error
    }
    return &shop, nil
}

// 获取所有宠物服务店铺
func (r *PetCareShopRepository) FindAllPetCareShops() ([]models.PetCareShop, error) {
    var shops []models.PetCareShop
    result := r.db.Find(&shops)
    return shops, result.Error
}

// 更新宠物服务店铺信息
func (r *PetCareShopRepository) UpdatePetCareShop(shop *models.PetCareShop) error {
    result := r.db.Save(shop)
    return result.Error
}

// 删除特定宠物服务店铺
func (r *PetCareShopRepository) DeletePetCareShop(shopID int) error {
    result := r.db.Delete(&models.PetCareShop{}, shopID)
    return result.Error
}
