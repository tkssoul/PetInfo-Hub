package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type PetFriendlySpotService struct {
    repo *repository.SpotRepository
}

// NewPetFriendlySpotService 创建一个新的 PetFriendlySpotService 实例
func NewPetFriendlySpotService(repo *repository.SpotRepository) *PetFriendlySpotService {
    return &PetFriendlySpotService{repo: repo}
}

// 创建宠物友好景点
func (s *PetFriendlySpotService) CreatePetFriendlySpot(spot *models.PetFriendlySpot) error {
    return s.repo.CreatePetFriendlySpot(spot)
}

// 获取特定景点
func (s *PetFriendlySpotService) GetPetFriendlySpotByID(spotID uint) (*models.PetFriendlySpot, error) {
    return s.repo.FindPetFriendlySpotByID(spotID)
}

// 获取所有景点
func (s *PetFriendlySpotService) GetAllPetFriendlySpots() ([]models.PetFriendlySpot, error) {
    return s.repo.FindAllPetFriendlySpots()
}

// 更新景点信息
func (s *PetFriendlySpotService) UpdatePetFriendlySpot(spot *models.PetFriendlySpot) error {
    existingSpot, err := s.repo.FindPetFriendlySpotByID(spot.SpotID)
    if err != nil {
        return err
    }
    if existingSpot == nil {
        return errors.New("景点不存在")
    }
    return s.repo.UpdatePetFriendlySpot(spot)
}

// 删除景点
func (s *PetFriendlySpotService) DeletePetFriendlySpot(spotID uint) error {
    existingSpot, err := s.repo.FindPetFriendlySpotByID(spotID)
    if err != nil {
        return err
    }
    if existingSpot == nil {
        return errors.New("景点不存在")
    }
    return s.repo.DeletePetFriendlySpot(spotID)
}
