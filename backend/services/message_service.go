package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type MessageService struct {
    repo *repository.MessageRepository
}

// NewMessageService 创建一个新的 MessageService 实例
func NewMessageService(repo *repository.MessageRepository) *MessageService {
    return &MessageService{repo: repo}
}

// CreateMessage 创建新的消息
func (s *MessageService) CreateMessage(message *models.Messages) error {
    return s.repo.CreateMessage(message)
}

// GetMessageByID 获取特定消息的详细信息
func (s *MessageService) GetMessageByID(messageID uint) (*models.Messages, error) {
    message, err := s.repo.FindMessageByID(messageID)
    if err != nil {
        return nil, err
    }
    if message == nil {
        return nil, errors.New("消息不存在")
    }
    return message, nil
}

// GetMessagesByUserID 获取特定用户的消息列表
func (s *MessageService) GetMessagesByUserID(userID uint) ([]models.Messages, error) {
    return s.repo.GetMessagesByUserID(userID)
}
