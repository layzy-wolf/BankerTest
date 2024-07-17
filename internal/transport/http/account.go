package http

import (
	"github.com/gin-gonic/gin"
	"github.com/layzy-wolf/BankerTest/internal/service"
)

type Amount struct {
	Amount float64 `json:"amount,omitempty"`
}

type bind struct {
	Id int `uri:"id" binding:"required"`
}

func CreateAccount(c *gin.Context, account *service.AccountService) {
	id := account.CreateAccount()
	c.JSON(200, gin.H{"id": id})
}

func Deposit(c *gin.Context, account *service.AccountService) {
	var a Amount
	var b bind
	if err := c.ShouldBindUri(&b); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := account.Deposit(b.Id, a.Amount)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"error": err})
}

func Withdraw(c *gin.Context, account *service.AccountService) {
	var a Amount
	var b bind
	if err := c.ShouldBindUri(&b); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	err := account.Withdraw(b.Id, a.Amount)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"error": err})
}

func GetBalance(c *gin.Context, account *service.AccountService) {
	var b bind
	if err := c.ShouldBindUri(&b); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	res := account.GetBalance(b.Id)

	c.JSON(200, gin.H{"balance": res})
}
