package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
	"strconv"
)

// PetFriendlySpotController 处理宠物友好景点相关的请求
type PetFriendlySpotController struct {
    spotService *services.PetFriendlySpotService
}

// NewPetFriendlySpotController 创建一个新的 PetFriendlySpotController 实例
func NewPetFriendlySpotController(spotService *services.PetFriendlySpotService) *PetFriendlySpotController {
    return &PetFriendlySpotController{spotService: spotService}
}

// CreatePetFriendlySpot 创建宠物友好景点
func (pc *PetFriendlySpotController) CreatePetFriendlySpot(c *gin.Context) {
    var spot models.PetFriendlySpot
    if err := c.ShouldBindJSON(&spot); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := pc.spotService.CreatePetFriendlySpot(&spot); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物友好景点创建成功"})
}

// GetPetFriendlySpotByID 获取特定宠物友好景点的详细信息
func (pc *PetFriendlySpotController) GetPetFriendlySpotByID(c *gin.Context) {
    spotIDStr := c.Param("spot_id")	
	guideIDUint, _ := strconv.ParseUint(spotIDStr, 10, 64)
	spotID := uint(guideIDUint) 
    spot, err := pc.spotService.GetPetFriendlySpotByID(spotID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, spot)
}

// GetAllPetFriendlySpots 获取所有宠物友好景点列表
func (pc *PetFriendlySpotController) GetAllPetFriendlySpots(c *gin.Context) {
    spots, err := pc.spotService.GetAllPetFriendlySpots()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, spots)
}

// UpdatePetFriendlySpot 更新特定宠物友好景点的信息
func (pc *PetFriendlySpotController) UpdatePetFriendlySpot(c *gin.Context) {
    var spot models.PetFriendlySpot
    if err := c.ShouldBindJSON(&spot); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pc.spotService.UpdatePetFriendlySpot(&spot)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物友好景点信息更新成功"})
}

// DeletePetFriendlySpot 删除特定宠物友好景点
func (pc *PetFriendlySpotController) DeletePetFriendlySpot(c *gin.Context) {
    spotIDStr := c.Param("spot_id")	
	guideIDUint, _ := strconv.ParseUint(spotIDStr, 10, 64)
	spotID := uint(guideIDUint) 
    err := pc.spotService.DeletePetFriendlySpot(spotID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物友好景点删除成功"})
}
