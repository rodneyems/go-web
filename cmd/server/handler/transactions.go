package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodneyems/go-web/internal/transactions"
)

type request struct {
	Currency string  `json:"currency" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Issuer   string  `json:"issuer" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

type requestPatch struct {
	Price    float64 `json:"price" binding:"required"`
	Issuer   string  `json:"issuer" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func (t Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != "1234567890" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "não autorizado",
			})
			return
		}
		t, _ := t.service.GetAll()
		c.JSON(200, t)
	}
}

func (t Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != "1234567890" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "não autorizado",
			})
			return
		}
		transaction := request{}
		err := c.ShouldBindJSON(&transaction)
		var ve validator.ValidationErrors
		fmt.Println(err)
		if errors.As(err, &ve) {
			for _, v := range ve {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("erro no campo: %v", v.Field()),
				})
				return
			}
		}
		t, _ := t.service.Save(transaction.Currency, transaction.Issuer, transaction.Receiver, transaction.Date, transaction.Price)
		c.JSON(http.StatusCreated, gin.H{
			"data": t,
		})
	}
}

func (t Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != "1234567890" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "não autorizado",
			})
			return
		}
		transaction := request{}
		err := c.ShouldBindJSON(&transaction)
		var ve validator.ValidationErrors
		fmt.Println(err)
		if errors.As(err, &ve) {
			for _, v := range ve {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("erro no campo: %v", v.Field()),
				})
				return
			}
		}
		id, err := strconv.Atoi(c.Param("id")) 
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "id inválido",
			})
			return
		}
		t, err := t.service.Update(id, transaction.Currency, transaction.Issuer, transaction.Receiver, transaction.Date, transaction.Price)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"data": t,
		})
	}
}

func (t Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != "1234567890" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "não autorizado",
			})
			return
		}
		id, err := strconv.Atoi(c.Param("id")) 
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "id inválido",
			})
			return
		}
		err = t.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "item removido",
		})
	}
}

func (t Transaction) UpdateFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != "1234567890" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "não autorizado",
			})
			return
		}
		transaction := requestPatch{}
		err := c.ShouldBindJSON(&transaction)
		var ve validator.ValidationErrors
		fmt.Println(err)
		if errors.As(err, &ve) {
			for _, v := range ve {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("erro no campo: %v", v.Field()),
				})
				return
			}
		}
		id, err := strconv.Atoi(c.Param("id")) 
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "id inválido",
			})
			return
		}
		t, err := t.service.UpdateFields(id, transaction.Issuer, transaction.Price)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"data": t,
		})
	}
}