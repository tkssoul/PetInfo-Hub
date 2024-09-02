package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type PetService struct {
    repo *repository.PetRepository
}

// NewPetService 创建一个新的 PetService 实例
func NewPetService(repo *repository.PetRepository) *PetService {
    return &PetService{repo: repo}
}

// 创建宠物
func (s *PetService) CreatePet(pet *models.Pets) error {
    return s.repo.CreatePet(pet)
}

// 通过ID查找宠物
func (s *PetService) FindPetByID(petID uint) (*models.Pets, error) {
    return s.repo.FindPetByID(petID)
}

// 更新宠物信息
func (s *PetService) UpdatePet(pet *models.Pets) error {
    existingPet, err := s.repo.FindPetByID(pet.Pet_ID)
    if err != nil {
        return err
    }
    if existingPet == nil {
        return errors.New("宠物不存在")
    }
    return s.repo.UpdatePet(pet)
}

// 删除宠物
func (s *PetService) DeletePet(petID uint) error {
    existingPet, err := s.repo.FindPetByID(petID)
    if err != nil {
        return err
    }
    if existingPet == nil {
        return errors.New("宠物不存在")
    }
    return s.repo.DeletePet(petID)
}

// 获取特定用户的宠物列表
func (s *PetService) GetPetsByUserID(userID uint) ([]models.Pets, error) {
    return s.repo.GetPetsByUserID(userID)
}
