package controllers

import (
	"fmt"
	"net/http"

	"backend-irvan-zulkarnain/config"
	"backend-irvan-zulkarnain/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	// Bind JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Ambil data produk
	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Cek stok
	if input.Quantity > product.Stock {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Kurangi stok
	product.Stock -= input.Quantity
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	fmt.Println("Harga:", product.Price)
	fmt.Println("Qty:", input.Quantity)

	totalPrice := input.Quantity * product.Price
	discount := 0
	if totalPrice > 50000 {
		discount = totalPrice / 10 // 10% diskon
		totalPrice = totalPrice - discount
	}

	shippingPrice := 2000
	if totalPrice > 15000 {
		shippingPrice = 0
	}
	totalPrice = totalPrice + shippingPrice
	// Buat order baru
	order := models.Order{
		ProductID:     product.ID,
		MerchantID:    product.MerchantID,
		Quantity:      input.Quantity,
		ShippingPrice: shippingPrice,
		TotalPrice:    totalPrice,
		Discount:      discount,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
