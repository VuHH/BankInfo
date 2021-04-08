package routers

import (
	"../controllers"
	"../middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {

	router.POST("/login", controllers.LoginController)

	api := router.Group("/api")

	api.Use(middleware.AuthorizeJWT())
	{
		api.GET("/users/:user_id/transactions/", controllers.GetTransactions)
		api.POST("/users/:user_id/transactions", controllers.CreateTransaction)
		//client.PATCH("/users/:id/transactions", controllers.Update)
		//client.DELETE("/users/:id/transactions", controllers.Delete)
	}

	return router
}
