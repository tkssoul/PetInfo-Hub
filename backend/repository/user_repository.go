package repository

import (
    "errors"
    "backend/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

// CreateUser 创建新用户
func (r *UserRepository) CreateUser(user *models.Users) error {
    result := r.db.Create(user)
    return result.Error
}

// FindUserByUsername 通过用户名查找用户
func (r *UserRepository) FindUserByUsername(username string) (*models.Users, error) {
    var user models.Users
    result := r.db.Where("username = ?", username).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("用户不存在")
        }
        return nil, result.Error
    }
    return &user, nil
}

// FindUserByID 通过用户ID查找用户
func (r *UserRepository) FindUserByID(userID int) (*models.Users, error) {
    var user models.Users
    result := r.db.First(&user, userID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("用户不存在")
        }
        return nil, result.Error
    }
    return &user, nil
}

// UpdateUser 更新用户信息
func (r *UserRepository) UpdateUser(user *models.Users) error {
    result := r.db.Save(user)
    return result.Error
}

// DeleteUser 删除用户
func (r *UserRepository) DeleteUser(userID int) error {
    result := r.db.Delete(&models.Users{}, userID)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("用户不存在")
    }
    return nil
}
