package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rodneyems/go-web/models"
)

func GetAllTransactions(c *gin.Context) {
	file, err := os.Open("transactions.json")
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message": "Erro no servidor, favor retornar mais tarde",
		})
		return
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message": "Erro no servidor, favor retornar mais tarde",
		})
		return
	}
	obj := []models.Transaction{}
	err = json.Unmarshal(byteValue, &obj)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
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
