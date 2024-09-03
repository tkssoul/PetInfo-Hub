package repository

import (
    "errors"
    "backend/models"
    "gorm.io/gorm"
)

type MessageRepository struct {
    db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
    return &MessageRepository{db: db}
}

// SendMessage 创建消息
func (r *MessageRepository) CreateMessage(message *models.Messages) error {
    result := r.db.Create(message)
    return result.Error
}

// GetMessagesByUserID 获取特定用户的消息列表
func (r *MessageRepository) GetMessagesByUserID(userID uint) ([]models.Messages, error) {
    var messages []models.Messages
    result := r.db.Where("receiver_id = ?", userID).Order("sent_at desc").Find(&messages)
    if result.Error != nil {
        return nil, result.Error
    }
    return messages, nil
}

// FindMessageByID 通过消息ID查找消息
func (r *MessageRepository) FindMessageByID(messageID uint) (*models.Messages, error) {
    var message models.Messages
    result := r.db.First(&message, messageID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("消息不存在")
        }
        return nil, result.Error
    }
    return &message, nil
}

// RevokeMessage 撤回消息
func (r *MessageRepository) RevokeMessage(messageID int) error {
    result := r.db.Where("message_id = ?", messageID).Delete(&models.Messages{})
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("消息不存在")
    }
    return nil
}
