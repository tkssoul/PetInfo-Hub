package repository

import (
    "errors"
    "backend/models"
    "gorm.io/gorm"
)

type RealNameRepository struct {
    db *gorm.DB
}

func NewRealNameRepository(db *gorm.DB) *RealNameRepository {
    return &RealNameRepository{db: db}
}

// CreateRealName 创建实名信息
func (r *RealNameRepository) CreateRealName(realName *models.RealName) error {
    result := r.db.Create(realName)
    return result.Error
}

// FindRealNameByUserID 通过用户ID查找实名信息
func (r *RealNameRepository) FindRealNameByUserID(userID int) (*models.RealName, error) {
    var realName models.RealName
    result := r.db.Where("user_id = ?", userID).First(&realName)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("实名信息不存在")
        }
        return nil, result.Error
    }
    return &realName, nil
}

// UpdateRealName 更新实名信息
func (r *RealNameRepository) UpdateRealName(realName *models.RealName) error {
    result := r.db.Save(realName)
    return result.Error
}

// DeleteRealName 删除实名信息
func (r *RealNameRepository) DeleteRealName(userID int) error {
    result := r.db.Where("user_id = ?", userID).Delete(&models.RealName{})
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("实名信息不存在")
    }
    return nil
}
