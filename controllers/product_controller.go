package controllers

import (
	"fmt"
	"net/http"

	"backend-irvan-zulkarnain/config"
	"backend-irvan-zulkarnain/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Ambil merchant_id dari JWT (contoh: disimpan di context)
	fmt.Println(c)
	merchantIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := merchantIDInterface.(uint)

	// Bind JSON body
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Overwrite merchant_id dari JWT, bukan dari body
	product.MerchantID = userID

	// Simpan ke database
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
