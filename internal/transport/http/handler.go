package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/layzy-wolf/BankerTest/internal/service"
)

func Handler(ctx *context.Context) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	account := service.NewAccountService(ctx)

	ac := r.Group("/accounts")
	{
		ac.POST("", func(c *gin.Context) {
			CreateAccount(c, account)
		})
		ac.POST("/:id/deposit", func(c *gin.Context) {
			Deposit(c, account)
		})
		ac.POST("/:id/withdraw", func(c *gin.Context) {
			Withdraw(c, account)
		})
		ac.GET("/:id/balance", func(c *gin.Context) {
			GetBalance(c, account)
		})
	}

	return r
}
