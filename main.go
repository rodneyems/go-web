package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodneyems/go-web/controllers"
)

func main() {
	router := gin.Default()
	group := router.Group("/transactions")
	{
		group.GET("/", controllers.GetAllTransactions)
		group.POST("/", controllers.AddTransactions)
	}

	router.Run()
}
