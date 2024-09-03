package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type UserService struct {
    repo *repository.UserRepository
    repo_realname *repository.RealNameRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

// 创建新用户
func (s *UserService) CreateUser(user *models.Users) error {    
    existingUser, err := s.repo.FindUserByUsername(user.Username)
    if err == nil && existingUser != nil {
        return errors.New("用户名已存在")
    }
    return s.repo.CreateUser(user)
}

// 通过用户名查找用户
func (s *UserService) FindUserByUsername(username string) (*models.Users, error) {
    return s.repo.FindUserByUsername(username)
}

// 通过用户ID查找用户
func (s *UserService) FindUserByID(userID uint) (*models.Users, error) {
    user, err := s.repo.FindUserByID(userID)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, errors.New("用户不存在")
    }
    return user, nil
}

// 更新用户信息
func (s *UserService) UpdateUser(user *models.Users) error {
    // 确保用户存在
    existingUser, err := s.repo.FindUserByID(user.User_ID)
    if err != nil {
        return err
    }
    if existingUser == nil {
        return errors.New("用户不存在")
    }
    return s.repo.UpdateUser(user)
}

// 删除用户
func (s *UserService) DeleteUser(userID uint) error {
    // 确保用户存在
    existingUser, err := s.repo.FindUserByID(userID)
    if err != nil {
        return err
    }
    if existingUser == nil {
        return errors.New("用户不存在")
    }
    return s.repo.DeleteUser(userID)
}

// 获取实名信息
func (s *UserService) GetRealNameInfo(userID uint) (*models.RealName, error) {
    return s.repo_realname.FindRealNameByUserID(userID)
}

// 创建实名信息
func (s *UserService) CreateRealNameInfo(realName *models.RealName) error {
    return s.repo_realname.CreateRealName(realName)
}

// 更新实名信息
func (s *UserService) UpdateRealNameInfo(realName *models.RealName) error {
    existingRealName, err := s.repo_realname.FindRealNameByUserID(realName.User_ID)
    if err != nil {
        return err
    }
    if existingRealName == nil {
        return errors.New("实名信息不存在")
    }
    return s.repo_realname.UpdateRealName(realName)
}

// 删除实名信息
func (s *UserService) DeleteRealNameInfo(userID uint) error {
    existingRealName, err := s.repo_realname.FindRealNameByUserID(userID)
    if err != nil {
        return err
    }
    if existingRealName == nil {
        return errors.New("实名信息不存在")
    }
    return s.repo_realname.DeleteRealName(userID)
}

// 获取所有用户信息
func (s *UserService) GetAllUsers() ([]models.Users, error) {
    return s.repo.GetAllUsers()
}