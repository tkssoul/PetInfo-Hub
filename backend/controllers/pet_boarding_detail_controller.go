package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
)

type PetBoardingDetailController struct {
    petBoardingDetailService *services.PetBoardingDetailService
}

func NewPetBoardingDetailController(petBoardingDetailService *services.PetBoardingDetailService) *PetBoardingDetailController {
    return &PetBoardingDetailController{petBoardingDetailService: petBoardingDetailService}
}

// GetBoardingDetailsBySitterID 获取特定寄养人的寄养信息列表
func (pb *PetBoardingDetailController) GetBoardingDetailsBySitterID(c *gin.Context) {
    sitterID, _ := strconv.Atoi(c.Param("sitterId"))

    boardingDetails, err := pb.petBoardingDetailService.GetPetBoardingDetailsBySitterID(sitterID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, boardingDetails)
}

// GetBoardingDetailByID 获取特定寄养信息的详细信息
func (pb *PetBoardingDetailController) GetBoardingDetailByID(c *gin.Context) {    
    boardIDStr := c.Param("boardingId")	
	boardIDUint, _ := strconv.ParseUint(boardIDStr, 10, 64)
	boardingID := uint(boardIDUint) 
    boardingDetail, err := pb.petBoardingDetailService.GetPetBoardingDetailByID(boardingID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, boardingDetail)
}

// CreateBoardingDetail 为特定寄养人创建寄养信息
func (pb *PetBoardingDetailController) CreateBoardingDetail(c *gin.Context) {
    var boardingDetail models.PetBoardingDetail
    if err := c.ShouldBindJSON(&boardingDetail); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pb.petBoardingDetailService.CreatePetBoardingDetail(&boardingDetail)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "寄养信息创建成功"})
}

// UpdateBoardingDetail 更新特定寄养信息
func (pb *PetBoardingDetailController) UpdateBoardingDetail(c *gin.Context) {
    var boardingDetail models.PetBoardingDetail
    if err := c.ShouldBindJSON(&boardingDetail); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pb.petBoardingDetailService.UpdatePetBoardingDetail(&boardingDetail)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "寄养信息更新成功"})
}

// DeleteBoardingDetail 删除特定寄养信息
func (pb *PetBoardingDetailController) DeleteBoardingDetail(c *gin.Context) {
    boardIDStr := c.Param("boardingId")	
	boardIDUint, _ := strconv.ParseUint(boardIDStr, 10, 64)
	boardingID := uint(boardIDUint) 
    err := pb.petBoardingDetailService.DeletePetBoardingDetail(boardingID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "寄养信息删除成功"})
}
