package controllers

import (
	"bankinfo.com/model/dto"
	"bankinfo.com/service"
	"github.com/gin-gonic/gin"
)

var transactionsList []dto.Transaction
var httpStatus int
var message interface{}

func GetTransactions(c *gin.Context) {
	httpStatus, message = service.GetTransactions(c.Param("user_id"), c.Query("account_id"))
	c.JSON(httpStatus, message)
}

func CreateTransaction(c *gin.Context) {
	var requestModel dto.RequestTransaction
	var err = c.ShouldBindJSON(&requestModel)
	httpStatus, message = service.CreateTransaction(c.Param("user_id"), requestModel, err)
	c.JSON(httpStatus, message)
}
