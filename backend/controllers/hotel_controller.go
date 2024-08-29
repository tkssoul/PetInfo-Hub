package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/services"
)

// 获取附近的酒店
func GetNearbyHotels(c *gin.Context) {
    hotels, err := services.GetNearbyHotels()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hotels)
}
