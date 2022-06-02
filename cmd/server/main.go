package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodneyems/go-web/cmd/server/handler"
	docs "github.com/rodneyems/go-web/docs"
	"github.com/rodneyems/go-web/internal/transactions"
	"github.com/rodneyems/go-web/pkg/store"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           MELI GO BOOTCAMP
// @version         1.0
// @description     API para gerencia de transações.

// @contact.name   Rodney Martins
// @contact.url    https://www.linkedin.com/in/rodney-martins
// @contact.email  rodney.martins@live.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	_ = godotenv.Load()
	store := store.New(store.FileType, "./pkg/store/file.json")
	repo := transactions.NewRepository(store)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:8080"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	v1 := router.Group("/api/v1")
	{
		group := v1.Group("/transactions")
		{
			group.GET("/", t.GetAll())
			group.POST("/", t.Store())
			group.PUT("/:id", t.Update())
			group.DELETE("/:id", t.Delete())
			group.PATCH("/:id", t.UpdateFields())
		}
	}
	router.Run()

}
