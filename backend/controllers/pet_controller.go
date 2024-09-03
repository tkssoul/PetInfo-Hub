package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
	"strconv"
)

type PetController struct {
    petService *services.PetService
}

func NewPetController(petService *services.PetService) *PetController {
    return &PetController{petService: petService}
}

// 创建宠物
func (pc *PetController) CreatePet(c *gin.Context) {
    var pet models.Pets
    if err := c.ShouldBindJSON(&pet); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := pc.petService.CreatePet(&pet); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物创建成功"})
}

// 通过ID获取宠物
func (pc *PetController) GetPetByID(c *gin.Context) {
    petIDStr := c.Param("pet_id")	
	userIDUint, _ := strconv.ParseUint(petIDStr, 10, 64)
	petID := uint(userIDUint) 
    pet, err := pc.petService.FindPetByID(petID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, pet)
}

// 更新宠物
func (pc *PetController) UpdatePet(c *gin.Context) {
    var pet models.Pets
    if err := c.ShouldBindJSON(&pet); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pc.petService.UpdatePet(&pet)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物信息更新成功"})
}

// 删除宠物
func (pc *PetController) DeletePet(c *gin.Context) {
    petIDStr := c.Param("pet_id")	
	userIDUint, err := strconv.ParseUint(petIDStr, 10, 64)
	petID := uint(userIDUint) 
    err1 := pc.petService.DeletePet(petID)
    if err1 != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "宠物删除成功"})
}

// 获取特定用户的宠物列表
func (pc *PetController) GetPetsByUserID(c *gin.Context) {
    userIdStr := c.Param("user_id")
	userIDUint, _ := strconv.ParseUint(userIdStr, 10, 64)
	userID := uint(userIDUint) 
    pets, err := pc.petService.GetPetsByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, pets)
}
