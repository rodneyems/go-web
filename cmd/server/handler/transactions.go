package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rodneyems/go-web/internal/transactions"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

type request struct {
	Currency string  `json:"currency" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Issuer   string  `json:"issuer" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

type requestPatch struct {
	Price  float64 `json:"price" binding:"required"`
	Issuer string  `json:"issuer" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

// Exibe Todas godoc
// @Summary      Exibe todas as transações
// @Description  Exibe todas as transacões cadastradas no sistema
// @Tags         Transactions
// @Produce      json
// @Param        Authorization  header    string  true  "token"
// @Success      200 {object} []transactions.transaction
// @Failure      400 {object} httputil.HTTPError
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /transactions [get]
func (t Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "não autorizado",
			})
			return
		}
		t, _ := t.service.GetAll()
		c.JSON(200, t)
	}
}

// Exibe Todas godoc
// @Summary      Cria uma nova transação no sistema
// @Description  Cria uma nova transação e retorna o objeto criado
// @Tags         Transactions
// @Receive	 	 json {object} request
// @Produce      json
// @Param        Authorization  header    string  true  "token"
// @Param 		 Transaction 	body 	  request    true "transaction"
// @Success      200 {object} transactions.transaction
// @Failure      400 {object} httputil.HTTPError
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /transactions [post]
func (t Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != os.Getenv("TOKEN") {
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

// Exibe Todas godoc
// @Summary      Atualiza uma nova transação no sistema
// @Description  Atualiza uma nova e retorna o objeto atualizado
// @Tags         Transactions
// @Receive	 	 json {object} request
// @Produce      json
// @Param        Authorization  header    string  true  "token"
// @Param		 id       path      int                   true  "transaction ID"
// @Param 		 Transaction 	body 	  request    true "transaction"
// @Success      200 {object} transactions.transaction
// @Failure      400 {object} httputil.HTTPError
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /transactions/{id} [put]
func (t Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != os.Getenv("TOKEN") {
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

// Exibe Todas godoc
// @Summary      Deleta uma transação do sistema
// @Description  Deleta uma transação do sistema
// @Tags         Transactions
// @Produce      json
// @Param        Authorization  header    string  true  "token"
// @Param		 id       path      int                   true  "transaction ID"
// @Success      200
// @Failure      400 {object} httputil.HTTPError
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /transactions/{id} [delete]
func (t Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != os.Getenv("TOKEN") {
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

// Exibe Todas godoc
// @Summary      Atualiza campos de uma transação no sistema
// @Description  Atualiza campos de uma transação e retorna o objeto atualizado
// @Tags         Transactions
// @Receive	 	 json {object} requestPatch
// @Produce      json
// @Param        Authorization  header    string  true  "token"
// @Param		 id       path      int                   true  "transaction ID"
// @Param 		 Transaction 	body 	  requestPatch    true "transaction"
// @Success      200 {object} transactions.transaction
// @Failure      400 {object} httputil.HTTPError
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /transactions/{id} [patch]
func (t Transaction) UpdateFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != os.Getenv("TOKEN") {
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
