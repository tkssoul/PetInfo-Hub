package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type UserRegistration struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type UserLogin struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func RegisterUser(userReg UserRegistration, userRepo *repository.UserRepository) error {
    // 检查用户是否已经存在
    existingUser, err := userRepo.FindUserByUsername(userReg.Username)
    if err == nil && existingUser != nil {
        return errors.New("用户已经存在")
    }

    // 创建新用户
    newUser := models.User{
        Username: userReg.Username,
        Password: userReg.Password, // In a real app, you'd hash this password
    }
    return userRepo.CreateUser(&newUser)
}

func LoginUser(userLogin UserLogin, userRepo *repository.UserRepository) (bool, error) {
    user, err := userRepo.FindUserByUsername(userLogin.Username)
    if err != nil {
        return false, err
    }

    if user.Password != userLogin.Password {
        return false, errors.New("密码错误")
    }

    return true, nil
}
