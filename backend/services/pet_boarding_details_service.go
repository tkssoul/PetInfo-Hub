package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type PetBoardingDetailService struct {
    repo *repository.PetBoardingDetailRepository
}

// NewPetBoardingDetailService 创建一个新的 PetBoardingDetailService 实例
func NewPetBoardingDetailService(repo *repository.PetBoardingDetailRepository) *PetBoardingDetailService {
    return &PetBoardingDetailService{repo: repo}
}

// 创建寄养信息
func (s *PetBoardingDetailService) CreatePetBoardingDetail(detail *models.PetBoardingDetail) error {
    return s.repo.CreatePetBoardingDetail(detail)
}

// 获取特定寄养信息
func (s *PetBoardingDetailService) GetPetBoardingDetailByID(boardingID int) (*models.PetBoardingDetail, error) {
    return s.repo.FindPetBoardingDetailByID(boardingID)
}

// 获取特定寄养人的所有寄养信息
func (s *PetBoardingDetailService) GetPetBoardingDetailsBySitterID(sitterID int) ([]models.PetBoardingDetail, error) {
    return s.repo.FindPetBoardingDetailsBySitterID(sitterID)
}

// 更新寄养信息
func (s *PetBoardingDetailService) UpdatePetBoardingDetail(detail *models.PetBoardingDetail) error {
    existingDetail, err := s.repo.FindPetBoardingDetailByID(detail.BoardingID)
    if err != nil {
        return err
    }
    if existingDetail == nil {
        return errors.New("寄养信息不存在")
    }
    return s.repo.UpdatePetBoardingDetail(detail)
}

// 删除寄养信息
func (s *PetBoardingDetailService) DeletePetBoardingDetail(boardingID int) error {
    existingDetail, err := s.repo.FindPetBoardingDetailByID(boardingID)
    if err != nil {
        return err
    }
    if existingDetail == nil {
        return errors.New("寄养信息不存在")
    }
    return s.repo.DeletePetBoardingDetail(boardingID)
}
