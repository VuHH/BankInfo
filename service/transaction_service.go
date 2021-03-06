package service

import (
	"bankinfo.com/model/dto"
	"bankinfo.com/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var transactionsList []dto.Transaction
var err string

func GetTransactions(userId string, accountId string) (int, interface{}) {
	if len(userId) > 0 {
		transactionsList, err = repository.FindTransactionsByUserIdAndAccountId(userId, accountId)
		if err == "" {
			if len(transactionsList) > 0 {
				return 200, transactionsList
			} else {
				return 200, gin.H{"messages": "No Records"}
			}

		} else {
			fmt.Println("err", err)
			return 502, gin.H{"messages": "System error"}
		}
	} else {
		return 501, gin.H{"messages": "User Id is required"}
	}

}

func CreateTransaction(userId string, requestModel dto.RequestTransaction, err error) (int, interface{}) {
	if len(userId) > 0 {
		if err == nil {
			if requestModel.TransactionType == "deposit" || requestModel.TransactionType == "withdraw" {
				if repository.SaveTransaction(requestModel) == "success" {
					return 200, gin.H{"messages": "Create Success!!"}
				} else {
					return 503, gin.H{"messages": "Create Failed"}
				}
			} else {
				return http.StatusBadRequest, gin.H{"error": "Transaction Type is deposit or withdraw"}
			}

		} else {
			return http.StatusBadRequest, gin.H{"error": err.Error()}
		}

	} else {
		return 501, gin.H{"messages": "User_Id is required"}
	}

}
