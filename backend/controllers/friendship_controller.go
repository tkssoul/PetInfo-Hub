package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "backend/services"
)

type FriendshipController struct {
    friendshipService *services.FriendshipService
}

func NewFriendshipController(friendshipService *services.FriendshipService) *FriendshipController {
    return &FriendshipController{friendshipService: friendshipService}
}

// AddFriend 添加好友
func (fc *FriendshipController) AddFriend(c *gin.Context) {
    userID, _ := strconv.Atoi(c.Param("userId"))
    friendID, _ := strconv.Atoi(c.Param("friendId"))

    err := fc.friendshipService.AddFriend(userID, friendID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "好友添加成功"})
}

// RemoveFriend 删除好友
func (fc *FriendshipController) RemoveFriend(c *gin.Context) {
    userID, _ := strconv.Atoi(c.Param("userId"))
    friendID, _ := strconv.Atoi(c.Param("friendId"))

    err := fc.friendshipService.RemoveFriend(userID, friendID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "好友删除成功"})
}

// GetFriendsByUserID 获取用户的好友列表
func (fc *FriendshipController) GetFriendsByUserID(c *gin.Context) {
    userID, _ := strconv.Atoi(c.Param("userId"))

    friendIDs, err := fc.friendshipService.GetFriendsByUserID(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"friends": friendIDs})
}
