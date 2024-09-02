package repository

import (
    "errors"
    "backend/models"
    "gorm.io/gorm"
)

type FriendshipRepository struct {
    db *gorm.DB
}

func NewFriendshipRepository(db *gorm.DB) *FriendshipRepository {
    return &FriendshipRepository{db: db}
}

// AddFriend 添加好友
func (r *FriendshipRepository) AddFriend(userID, friendID int) error {
    // 确保一个用户不能与自己成为好友
    if userID == friendID {
        return errors.New("不能添加自己为好友")
    }

    // 创建好友关系
    friendship := models.Friendship{
        User_ID:   userID,
        Friend_ID: friendID,
    }

    result := r.db.Create(&friendship)
    if result.Error != nil {
        // 如果是唯一约束错误，说明好友关系已经存在
        if result.Error.Error() == "UNIQUE constraint failed: Friendships.user_id, Friendships.friend_id" {
            return errors.New("好友关系已存在")
        }
        return result.Error
    }
    return nil
}

// RemoveFriend 删除好友
func (r *FriendshipRepository) RemoveFriend(userID, friendID int) error {
    // 确保一个用户不能与自己成为好友
    if userID == friendID {
        return errors.New("不能删除自己")
    }

    // 删除好友关系
    result := r.db.Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&models.Friendship{})
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("好友关系不存在")
    }
    return nil
}

// GetFriendsByUserID 获取特定用户的好友列表
func (r *FriendshipRepository) GetFriendsByUserID(userID int) ([]int, error) {
    var friendships []models.Friendship
    result := r.db.Where("user_id = ?", userID).Find(&friendships)
    if result.Error != nil {
        return nil, result.Error
    }

    // 提取好友ID
    friendIDs := make([]int, 0, len(friendships))
    for _, friendship := range friendships {
        friendIDs = append(friendIDs, friendship.Friend_ID)
    }
    return friendIDs, nil
}

// FindUserByID 通过用户ID查找用户
func (r *FriendshipRepository) FindUserByID(userID int) (*models.Users, error) {
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
