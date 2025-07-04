package controllers

import (
	"net/http"

	"backend-irvan-zulkarnain/config"
	"backend-irvan-zulkarnain/models"
	"backend-irvan-zulkarnain/utils"

	"github.com/gin-gonic/gin"
)

func CreateMerchant(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(merchant.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	merchant.Password = hashedPassword

	if err := config.DB.Create(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create merchant"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func GetMerchants(c *gin.Context) {
	var merchants []models.Merchant
	if err := config.DB.Find(&merchants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get merchants"})
		return
	}

	c.JSON(http.StatusOK, merchants)
}
