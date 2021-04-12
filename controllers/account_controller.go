package controllers

import (
	"bankinfo.com/model/dto"
	"bankinfo.com/service"
	"github.com/gin-gonic/gin"
)

func UpdateAccounts(c *gin.Context) {
	var updateAccount dto.UpdateAccount
	var err = c.ShouldBindJSON(&updateAccount)
	httpStatus, message = service.UpdateNameAccount(c.Param("account_id"), updateAccount, err)
	c.JSON(httpStatus, message)
}

func DeleteAccount(c *gin.Context) {
	var updateAccount dto.UpdateAccount
	var err = c.ShouldBindJSON(&updateAccount)
	httpStatus, message = service.UpdateStatusActiveAccount(c.Param("account_id"), updateAccount, err)
	c.JSON(httpStatus, message)
}
