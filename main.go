package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodneyems/go-web/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/greetings", )
	router.GET("/transactions", controllers.GetAllTransactions)
	router.Run()
}
