package service

import (
	"github.com/gin-gonic/gin"
)

func GetTransactionsUser(c *gin.Context) {
	var userId = c.Param("user_id")
	var accountId = c.Query("account_id")

	if len() > 0 {

	} else {
		c.JSON(501, gin.H{
			"messages": "User Id is required",
		})
	}

}

//func CreateTransaction(c * gin.Context) {
//	db := config.DBConn()
//
//	if len(c.Param("user_id")) > 0 {
//		var json model.RequestTransaction
//		if err := c.ShouldBindJSON(&json); err == nil {
//			insTran, err := db.Prepare(
//				"INSERT INTO transactions (id_account, amount, transaction_type) " +
//					"VALUES(?, ?, ?)")
//			if err != nil {
//				c.JSON(500, gin.H{
//					"messages" : err,
//				})
//			}
//			insTran.Exec(json.AccountId,json.Amount,json.TransactionType)
//			c.JSON(200, gin.H{
//				"messages": "inserted",
//			})
//		} else {
//			c.JSON(500, gin.H{"error": err.Error()})
//		}
//
//	} else {
//		c.JSON(500, gin.H{
//			"messages" : "User_Id is required",
//		})
//	}
//
//	defer db.Close()
//}
