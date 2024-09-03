package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type PetSitterService struct {
    repo *repository.PetSitterRepository
}

// NewPetSitterService 创建一个新的 PetSitterService 实例
func NewPetSitterService(repo *repository.PetSitterRepository) *PetSitterService {
    return &PetSitterService{repo: repo}
}

// 创建寄养人
func (s *PetSitterService) CreatePetSitter(sitter *models.PetSitter) error {
    return s.repo.CreatePetSitter(sitter)
}

// 获取特定寄养人
func (s *PetSitterService) GetPetSitterByID(sitterID uint) (*models.PetSitter, error) {
    return s.repo.FindPetSitterByID(sitterID)
}

// 获取所有寄养人
func (s *PetSitterService) GetAllPetSitters() ([]models.PetSitter, error) {
    return s.repo.FindAllPetSitters()
}

// 更新寄养人信息
func (s *PetSitterService) UpdatePetSitter(sitter *models.PetSitter) error {
    existingSitter, err := s.repo.FindPetSitterByID(sitter.ID)
    if err != nil {
        return err
    }
    if existingSitter == nil {
        return errors.New("寄养人不存在")
    }
    return s.repo.UpdatePetSitter(sitter)
}

// 删除寄养人
func (s *PetSitterService) DeletePetSitter(sitterID uint) error {
    existingSitter, err := s.repo.FindPetSitterByID(sitterID)
    if err != nil {
        return err
    }
    if existingSitter == nil {
        return errors.New("寄养人不存在")
    }
    return s.repo.DeletePetSitter(sitterID)
}
