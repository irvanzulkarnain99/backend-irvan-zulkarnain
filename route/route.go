package route

import (
	controllers "backend-irvan-zulkarnain/controllers"
	middleware "backend-irvan-zulkarnain/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Sementara bisa kosong atau isi default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	router.POST("/login", controllers.Login)

	merchant := router.Group("/merchants")
	{
		merchant.POST("/createMerchant", controllers.CreateMerchant)
		merchant.GET("/getMerchant", middleware.AuthMiddleware(), controllers.GetMerchants)
	}

	product := router.Group("/products")
	{
		product.POST("/createProduct", middleware.AuthMiddleware(), controllers.CreateProduct)
		product.GET("/getProduct", middleware.AuthMiddleware(), controllers.GetProducts)
	}

	order := router.Group("/orders")
	{
		order.POST("/createOrder", middleware.AuthMiddleware(), controllers.CreateOrder)
		order.GET("/getOrder", middleware.AuthMiddleware(), controllers.GetOrders)
	}
}
