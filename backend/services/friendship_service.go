package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type FriendshipService struct {
    repo *repository.FriendshipRepository
}

func NewFriendshipService(repo *repository.FriendshipRepository) *FriendshipService {
    return &FriendshipService{repo: repo}
}

// 添加好友
func (s *FriendshipService) AddFriend(userID, friendID int) error {
    // 检查输入的用户ID是否有效
    if userID <= 0 || friendID <= 0 {
        return errors.New("无效的用户ID")
    }
    // 调用仓库方法创建好友关系
    return s.repo.AddFriend(userID, friendID)
}

// 删除好友
func (s *FriendshipService) RemoveFriend(userID, friendID int) error {
    // 检查输入的用户ID是否有效
    if userID <= 0 || friendID <= 0 {
        return errors.New("无效的用户ID")
    }
    // 调用仓库方法删除好友关系
    return s.repo.RemoveFriend(userID, friendID)
}

// 获取特定用户的好友列表
func (s *FriendshipService) GetFriendsByUserID(userID int) ([]models.Users, error) {
    // 获取好友ID列表
    friendIDs, err := s.repo.GetFriendsByUserID(userID)
    if err != nil {
        return nil, err
    }

    // 根据好友ID列表获取好友详细信息
    friends := make([]models.Users, 0, len(friendIDs))
    for _, friendID := range friendIDs {
        friend, err := s.repo.FindUserByID(friendID)
        if err != nil {
            return nil, err
        }
        if friend != nil {
            friends = append(friends, *friend)
        }
    }
    return friends, nil
}
