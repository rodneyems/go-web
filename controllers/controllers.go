package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodneyems/go-web/models"
)

func GetAllTransactions(c *gin.Context) {
	c.JSON(200, models.Transactions)
}

func Greetings(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Olá Rodney",
	})
}

func AddTransactions(c *gin.Context) {
	if c.GetHeader("token") != "1234567890" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "não autorizado",
		})
		return
	}
	transaction := models.Transaction{}
	err := c.ShouldBindJSON(&transaction)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, v := range ve {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("erro no campo: %v", v.Field()),
			})
			return
		}
	}
	models.LastId++
	transaction.Id = models.LastId
	models.Transactions = append(models.Transactions, transaction)
	c.JSON(http.StatusCreated, gin.H{
		"data": transaction,
	})
}
