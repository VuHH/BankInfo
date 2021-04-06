package service

import (
	"../model/dto"
	"../repository"
	"github.com/gin-gonic/gin"
)

var transactionsList []dto.Transaction
var err string

func GetTransactions(c *gin.Context) {
	var userId = c.Param("user_id")
	var accountId = c.Query("account_id")

	if len(userId) > 0 {
		transactionsList, err = repository.FindTransactionsByUserIdAndAccountId(userId, accountId)
		if len(transactionsList) > 0 {
			c.JSON(200, transactionsList)
		} else {
			// TODO: log error
			c.JSON(502, gin.H{
				"messages": "System error",
			})
		}
	} else {
		c.JSON(501, gin.H{
			"messages": "User Id is required",
		})
	}

}

func CreateTransaction(c *gin.Context) {
	var userId = c.Param("user_id")
	if len(userId) > 0 {
		var requestModel dto.RequestTransaction
		if err := c.ShouldBindJSON(&requestModel); err == nil {
			if repository.SaveTransaction(requestModel) == "success" {
				c.JSON(200, gin.H{
					"messages": "Create Success!!",
				})
			} else {
				c.JSON(503, gin.H{
					"messages": "Create Failed",
				})
			}

		} else {
			//TODO log error err.Error()
			c.JSON(500, gin.H{"error": "System error"})
		}

	} else {
		c.JSON(500, gin.H{
			"messages": "User_Id is required",
		})
	}

}
