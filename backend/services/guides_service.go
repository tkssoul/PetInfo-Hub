package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type GuideService struct {
    repo *repository.GuideRepository
}

// NewGuideService 创建一个新的 GuideService 实例
func NewGuideService(repo *repository.GuideRepository) *GuideService {
    return &GuideService{repo: repo}
}

// 创建攻略
func (s *GuideService) CreateGuide(guide models.Guide) (models.Guide,error) {
    return s.repo.CreateGuide(guide)
}

// 获取特定攻略
func (s *GuideService) GetGuideByID(guideID int) (*models.Guide, error) {
    return s.repo.GetGuideByID(guideID)
}

// 获取所有攻略
func (s *GuideService) GetAllGuides() ([]models.Guide, error) {
    return s.repo.GetAllGuides()
}

// 更新攻略信息
func (s *GuideService) UpdateGuide(guide models.Guide) error {
    existingGuide, err := s.repo.GetGuideByID(guide.Guide_ID)
    if err != nil {
        return err
    }
    if existingGuide == nil {
        return errors.New("攻略不存在")
    }
    return s.repo.UpdateGuide(guide)
}

// 删除攻略
func (s *GuideService) DeleteGuide(guideID int) error {
    existingGuide, err := s.repo.GetGuideByID(guideID)
    if err != nil {
        return err
    }
    if existingGuide == nil {
        return errors.New("攻略不存在")
    }
    return s.repo.DeleteGuide(guideID)
}
