package controllers

import (
	"fmt"
	"net/http"

	"backend-irvan-zulkarnain/config"
	"backend-irvan-zulkarnain/models"
	"backend-irvan-zulkarnain/utils"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	fmt.Println("BindJSON customer:", customer.Password)
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(customer.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	customer.Password = hashedPassword

	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	if err := config.DB.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}
