package handlers

import (
	"../service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.GET("/users/:user_id/transactions/", service.GetTransactions)
		client.POST("/users/:user_id/transactions", service.CreateTransaction)
		//client.PATCH("/users/:id/transactions", controllers.Update)
		//client.DELETE("/users/:id/transactions", controllers.Delete)
	}

	return r
}
