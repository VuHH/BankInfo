package routers

import (
	"bankinfo.com/controllers"
	"bankinfo.com/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/login", controllers.LoginController)

	api := router.Group("/api")

	api.Use(middleware.AuthorizeJWT())
	{
		api.GET("/users/:user_id/transactions", controllers.GetTransactions)
		api.POST("/users/:user_id/transactions", controllers.CreateTransaction)
		api.POST("/accounts/:account_id/update", controllers.UpdateAccounts)
		api.POST("/accounts/:account_id/remove", controllers.DeleteAccount)
		//client.DELETE("/users/:id/transactions", controllers.Delete)
	}

	return router
}
