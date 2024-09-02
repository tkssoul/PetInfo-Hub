package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
)

type GuideController struct {
    guideService *services.GuideService
}

func NewGuideController(guideService *services.GuideService) *GuideController {
    return &GuideController{guideService: guideService}
}

// CreateGuide 创建攻略
func (gc *GuideController) CreateGuide(c *gin.Context) {
    var guide models.Guide
    if err := c.ShouldBindJSON(&guide); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdGuide, err := gc.guideService.CreateGuide(guide)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"guide": createdGuide})
}

// GetGuideByID 通过ID获取攻略
func (gc *GuideController) GetGuideByID(guideID uint,c *gin.Context) {        
    guide, err := gc.guideService.GetGuideByID(guideID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, guide)
}

// GetAllGuides 获取所有攻略
func (gc *GuideController) GetAllGuides(c *gin.Context) {
    guides, err := gc.guideService.GetAllGuides()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, guides)
}

// UpdateGuide 更新攻略
func (gc *GuideController) UpdateGuide(c *gin.Context) {
    var guide models.Guide
    if err := c.ShouldBindJSON(&guide); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := gc.guideService.UpdateGuide(guide)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "攻略更新成功"})
}

// DeleteGuide 删除攻略
func (gc *GuideController) DeleteGuide(guideID uint,c *gin.Context) {
    err := gc.guideService.DeleteGuide(guideID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "攻略删除成功"})
}
