package routers

import (
	"bankinfo.com/controllers"
	"bankinfo.com/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

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
