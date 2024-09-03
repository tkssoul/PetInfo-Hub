package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
	"strconv"
)

// PetCareShopController 处理宠物服务店铺相关的请求
type PetCareShopController struct {
    shopService *services.PetCareShopService
}

// NewPetCareShopController 创建一个新的 PetCareShopController 实例
func NewPetCareShopController(shopService *services.PetCareShopService) *PetCareShopController {
    return &PetCareShopController{shopService: shopService}
}

// CreatePetCareShop 创建宠物服务店铺
func (pc *PetCareShopController) CreatePetCareShop(c *gin.Context) {
    var shop models.PetCareShop
    if err := c.ShouldBindJSON(&shop); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := pc.shopService.CreatePetCareShop(&shop); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物服务店铺创建成功"})
}

// GetPetCareShopByID 获取特定宠物服务店铺的详细信息
func (pc *PetCareShopController) GetPetCareShopByID(c *gin.Context) {
    shopIDStr := c.Param("shop_id")	
	shopIDUint, _ := strconv.ParseUint(shopIDStr, 10, 64)
	shopID := uint(shopIDUint) 
    shop, err := pc.shopService.GetPetCareShopByID(shopID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, shop)
}

// GetAllPetCareShops 获取所有宠物服务店铺列表
func (pc *PetCareShopController) GetAllPetCareShops(c *gin.Context) {
    shops, err := pc.shopService.GetAllPetCareShops()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, shops)
}

// UpdatePetCareShop 更新特定宠物服务店铺的信息
func (pc *PetCareShopController) UpdatePetCareShop(c *gin.Context) {
    var shop models.PetCareShop
    if err := c.ShouldBindJSON(&shop); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pc.shopService.UpdatePetCareShop(&shop)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物服务店铺信息更新成功"})
}

// DeletePetCareShop 删除特定宠物服务店铺
func (pc *PetCareShopController) DeletePetCareShop(c *gin.Context) {
    shopIDStr := c.Param("shop_id")	
	shopIDUint, _ := strconv.ParseUint(shopIDStr, 10, 64)
	shopID := uint(shopIDUint) 
    err := pc.shopService.DeletePetCareShop(shopID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物服务店铺删除成功"})
}
