package service

import (
	"../model/dto"
	"../repository"
	"github.com/gin-gonic/gin"
)

var transactionsList []dto.Transaction
var err string

func GetTransactions(userId string, accountId string) (int, interface{}) {
	if len(userId) > 0 {
		transactionsList, err = repository.FindTransactionsByUserIdAndAccountId(userId, accountId)
		if err == "" {
			return 200, transactionsList
		} else {
			//log.Fatal(err)
			return 502, gin.H{"messages": "System error"}
		}
	} else {
		return 501, gin.H{"messages": "User Id is required"}
	}

}

func CreateTransaction(userId string, requestModel dto.RequestTransaction, err error) (int, interface{}) {
	if len(userId) > 0 {
		if err == nil {
			if repository.SaveTransaction(requestModel) == "success" {
				return 200, gin.H{"messages": "Create Success!!"}
			} else {
				return 503, gin.H{"messages": "Create Failed"}
			}

		} else {
			//log.Fatal(err)
			return 500, gin.H{"error": "System error"}
		}

	} else {
		return 501, gin.H{"messages": "User_Id is required"}
	}

}
