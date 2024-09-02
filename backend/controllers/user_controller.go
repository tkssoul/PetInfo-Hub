package controllers

import (
    "net/http"    
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{userService: userService}
}

// CreateUser 创建用户
func (uc *UserController) CreateUser(c *gin.Context) {
    var user models.Users
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := uc.userService.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "用户创建成功"})
}

// GetUserByID 通过ID获取用户
func (uc *UserController) GetUserByID(userID uint,c *gin.Context) {
    user, err := uc.userService.FindUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// UpdateUser 更新用户
func (uc *UserController) UpdateUser(c *gin.Context) {
    var user models.Users
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := uc.userService.UpdateUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}

// DeleteUser 删除用户
func (uc *UserController) DeleteUser(userID uint,c *gin.Context) {
    err := uc.userService.DeleteUser(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}
