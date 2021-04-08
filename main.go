package main

import (
	"./routers"
	"github.com/gin-gonic/gin"
)

//func setupRouter() *gin.Engine {
//	r := gin.Default()
//	r.Static("/public", "./public")
//
//	client := r.Group("/api")
//	{
//		client.GET("/users/:user_id/transactions/", handlers.GetTransactionsUser)
//		client.POST("/users/:user_id/transactions", handlers.CreateTransaction)
//		//client.PATCH("/users/:id/transactions", controllers.Update)
//		//client.DELETE("/users/:id/transactions", controllers.Delete)
//	}
//
//	return r
//}

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	r := routers.SetupRouter(router)
	r.Run(":8080")
}
