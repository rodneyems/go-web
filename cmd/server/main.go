package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodneyems/go-web/cmd/server/handler"
	"github.com/rodneyems/go-web/internal/transactions"
	"github.com/rodneyems/go-web/pkg/store"
)

func main() {
	_ = godotenv.Load()
	store := store.New(store.FileType, "./pkg/store/file.json")
	repo := transactions.NewRepository(store)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	router := gin.Default()
	group := router.Group("/transactions")
	{
		group.GET("/", t.GetAll())
		group.POST("/", t.Store())
		group.PUT("/:id", t.Update())
		group.DELETE("/:id", t.Delete())
		group.PATCH("/:id", t.UpdateFields())
	}

	router.Run()

}
