package service

import (
	"bankinfo.com/model/dto"
	"bankinfo.com/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpdateNameAccount(accountId string, updateAccount dto.UpdateAccount, err error) (int, interface{}) {
	if len(accountId) > 0 {
		if err != nil {
			fmt.Println("err", err)
			return 500, gin.H{"messages": "System error"}
		} else {
			if updateAccount.NameBank == "" {
				return 501, gin.H{"messages": "Name Bank is required"}
			} else {
				if repository.UpdateNameAccount(updateAccount.NameBank, accountId) == "success" {
					return 200, gin.H{"messages": "Update Success!!"}
				} else {
					return 503, gin.H{"messages": "Update Failed"}
				}
			}
		}
	} else {
		return 501, gin.H{"messages": "Account Id is required"}
	}

}

func UpdateStatusActiveAccount(accountId string, updateAccount dto.UpdateAccount, err error) (int, interface{}) {
	if len(accountId) > 0 {
		if err != nil {
			fmt.Println("err", err)
			return 500, gin.H{"messages": "System error"}
		} else {
			if repository.UpdateStatusAccount(updateAccount.IsActive, accountId) == "success" {
				return 200, gin.H{"messages": "Update Success!!"}
			} else {
				return 503, gin.H{"messages": "Update Failed"}
			}
		}

	} else {
		return 501, gin.H{"messages": "Account Id is required"}
	}

}
