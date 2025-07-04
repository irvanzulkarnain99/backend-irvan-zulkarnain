package main

import (
	"backend-irvan-zulkarnain/config"
	"backend-irvan-zulkarnain/route"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	route.SetupRoutes(r)

	r.Run() // default di :8080
}
