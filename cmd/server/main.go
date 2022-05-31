package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodneyems/go-web/cmd/server/handler"
	"github.com/rodneyems/go-web/internal/transactions"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	router := gin.Default()
	group := router.Group("/transactions")
	{
		group.GET("/", t.GetAll())
		group.POST("/", t.Store())
		group.PUT("/:id", t.Update())
		group.DELETE("/:id", t.Delete())
	}

	router.Run()

}
