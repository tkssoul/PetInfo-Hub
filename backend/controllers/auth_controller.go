package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/services"
    "backend/repository"
    "backend/models"
)

func Register(c *gin.Context) {
    var userReg services.UserRegistration
    if err := c.BindJSON(&userReg); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    userRepo := repository.NewUserRepository(models.DB)
    err := services.RegisterUser(userReg, userRepo)
    if err != nil {
        if err.Error() == "用户已经存在" {
            c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
    var userLogin services.UserLogin
    if err := c.BindJSON(&userLogin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    userRepo := repository.NewUserRepository(models.DB)
    valid, err := services.LoginUser(userLogin, userRepo)
    if err != nil {
        if err.Error() == "用户不存在" {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        } else if err.Error() == "密码错误" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }       
        return
    }

    if valid {
        c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    }
}
