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

func (r *UserRepository) CreateUser(user *models.User) error {
    result := r.db.Create(user)
    return result.Error
}

func (r *UserRepository) FindUserByUsername(username string) (*models.User, error) {
    var user models.User
    result := r.db.Where("username = ?", username).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("用户不存在")
        }
        return nil, result.Error
    }
    return &user, nil
}
