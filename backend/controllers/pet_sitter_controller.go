package controllers

import (
    "net/http"    
    "github.com/gin-gonic/gin"    
    "backend/services"
	"strconv"
)

type PetSitterController struct {
    petSitterService *services.PetSitterService
}

func NewPetSitterController(petSitterService *services.PetSitterService) *PetSitterController {
    return &PetSitterController{petSitterService: petSitterService}
}

// GetAllPetSitters 获取所有寄养人
func (psc *PetSitterController) GetAllPetSitters(c *gin.Context) {
    petSitters, err := psc.petSitterService.GetAllPetSitters()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, petSitters)
}

// GetPetSitterByID 通过ID获取
func (psc *PetSitterController) GetPetSitterByID(c *gin.Context) {
	petSitterIDStr := c.Param("guide_id")	
	petSitterIDUint, _ := strconv.ParseUint(petSitterIDStr, 10, 64)
	petSitterID := uint(petSitterIDUint) 
	petSitter, err := psc.petSitterService.GetPetSitterByID(petSitterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, petSitter)
}