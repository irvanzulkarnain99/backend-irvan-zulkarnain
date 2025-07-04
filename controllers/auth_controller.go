package controllers

import (
	"fmt"
	"net/http"

	"backend-irvan-zulkarnain/config"
	"backend-irvan-zulkarnain/models"
	"backend-irvan-zulkarnain/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var merchant models.Merchant
	if err := config.DB.Where("email = ?", input.Email).First(&merchant).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found"})
		return
	}
	// Debug
	fmt.Println("Input Password:", input.Password)
	fmt.Println("DB Hash:", merchant.Password)

	// Cek password
	if !utils.CheckPasswordHash(input.Password, merchant.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := utils.GenerateJWT(merchant.ID, "merchant")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
