package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodneyems/go-web/models"
)

func GetAllTransactions(c *gin.Context) {
	file, err := os.Open("transactions.json")
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Erro no servidor, favor retornar mais tarde",
		})
		return
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Erro no servidor, favor retornar mais tarde",
		})
		return
	}
	obj := []models.Transaction{}
	err = json.Unmarshal(byteValue, &obj)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Erro no servidor, favor retornar mais tarde",
		})
		return
	}
	c.JSON(200, obj)
}

func Greetings(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Olá Rodney",
	})
}

func GetTransactionById(c *gin.Context) {
	file, err := os.Open("transactions.json")
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Erro no servidor, favor retornar mais tarde",
		})
		return
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Erro no servidor, favor retornar mais tarde",
		})
		return
	}
	obj := []models.Transaction{}
	err = json.Unmarshal(byteValue, &obj)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "erro no servidor, favor retornar mais tarde",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"status":  400,
			"message": "id inválido",
		})
		return
	}
	for _, v := range obj {
		if v.Id == id {
			c.JSON(200, v)
			return
		}
	}
	c.JSON(200, gin.H{
		"status":  404,
		"message": "id nao encontrado",
	})
}
