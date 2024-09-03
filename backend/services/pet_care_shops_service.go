package services	

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type PetCareShopService struct {
    repo *repository.PetCareShopRepository
}

// NewPetCareShopService 创建一个新的 PetCareShopService 实例
func NewPetCareShopService(repo *repository.PetCareShopRepository) *PetCareShopService {
    return &PetCareShopService{repo: repo}
}

// 创建宠物服务店铺
func (s *PetCareShopService) CreatePetCareShop(shop *models.PetCareShop) error {
    return s.repo.CreatePetCareShop(shop)
}

// 获取特定店铺
func (s *PetCareShopService) GetPetCareShopByID(shopID uint) (*models.PetCareShop, error) {
    return s.repo.FindPetCareShopByID(shopID)
}

// 获取所有店铺
func (s *PetCareShopService) GetAllPetCareShops() ([]models.PetCareShop, error) {
    return s.repo.FindAllPetCareShops()
}

// 更新店铺信息
func (s *PetCareShopService) UpdatePetCareShop(shop *models.PetCareShop) error {
    existingShop, err := s.repo.FindPetCareShopByID(shop.Shop_ID)
    if err != nil {
        return err
    }
    if existingShop == nil {
        return errors.New("店铺不存在")
    }
    return s.repo.UpdatePetCareShop(shop)
}

// 删除店铺
func (s *PetCareShopService) DeletePetCareShop(shopID uint) error {
    existingShop, err := s.repo.FindPetCareShopByID(shopID)
    if err != nil {
        return err
    }
    if existingShop == nil {
        return errors.New("店铺不存在")
    }
    return s.repo.DeletePetCareShop(shopID)
}
