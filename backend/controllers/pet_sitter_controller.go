package controllers

import (
    "net/http"    
    "github.com/gin-gonic/gin"    
    "backend/services"
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
