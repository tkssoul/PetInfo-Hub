package controllers

import (
    "net/http"    
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
	"strconv"
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
    user.Username = c.PostForm("username")
    user.Password = c.PostForm("password")
    if err := c.ShouldBind(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"传入的数据格式错误,error:": err.Error()})
        return
    }

    if err := uc.userService.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"创建用户失败,error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "用户创建成功"})
}

// GetUserByID 通过ID获取用户
func (uc *UserController) GetUserByID(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userIDUint, _ := strconv.ParseUint(userIdStr, 10, 64)
	userID := uint(userIDUint) 
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
func (uc *UserController) DeleteUser(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userIDUint, _ := strconv.ParseUint(userIdStr, 10, 64)
	userID := uint(userIDUint) 
    err := uc.userService.DeleteUser(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }    
    c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

// GetRealNameInfo 获取用户实名信息
func (uc *UserController) GetRealNameInfo(c *gin.Context) {
    userIdStr := c.Param("user_id")
	userIDUint, _ := strconv.ParseUint(userIdStr, 10, 64)
	userID := uint(userIDUint) 
    realNameInfo, err := uc.userService.GetRealNameInfo(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, realNameInfo)
}

// CreateRealNameInfo 为用户创建实名信息
func (uc *UserController) CreateRealNameInfo(c *gin.Context) {
    var realNameInfo models.RealName
    if err := c.ShouldBindJSON(&realNameInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    if err := uc.userService.CreateRealNameInfo(&realNameInfo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "实名信息创建成功"})
}

// UpdateRealNameInfo 更新用户实名信息
func (uc *UserController) UpdateRealNameInfo(c *gin.Context) {
    var realNameInfo models.RealName
    if err := c.ShouldBindJSON(&realNameInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    if err := uc.userService.UpdateRealNameInfo(&realNameInfo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "实名信息更新成功"})
}

// DeleteRealNameInfo 删除用户实名信息
func (uc *UserController) DeleteRealNameInfo(c *gin.Context) {
    userIdStr := c.Param("user_id")
	userIDUint, _ := strconv.ParseUint(userIdStr, 10, 64)
	userID := uint(userIDUint) 
    if err := uc.userService.DeleteRealNameInfo(userID); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "实名信息删除成功"})
}

// GetAllUsers 获取所有用户列表
func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.userService.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}
