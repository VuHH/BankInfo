package handlers

import (
	"../middleware"
	"../service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//r := gin.Default()
	//r.Static("/public", "./public")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	//r.POST("/login", func(ctx *gin.Context) {
	//	token := loginController.Login(ctx)
	//	if token != "" {
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"token": token,
	//		})
	//	} else {
	//		ctx.JSON(http.StatusUnauthorized, nil)
	//	}
	//})

	router.POST("/login", service.LoginService)

	api := router.Group("/api")

	api.Use(middleware.AuthorizeJWT())
	{
		api.GET("/users/:user_id/transactions/", service.GetTransactions)
		api.POST("/users/:user_id/transactions", service.CreateTransaction)
		//client.PATCH("/users/:id/transactions", controllers.Update)
		//client.DELETE("/users/:id/transactions", controllers.Delete)
	}

	return router
}
