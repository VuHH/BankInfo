package repository

import (
	"../config"
	"../model/dto"
	_ "database/sql"
	"time"
)

func FindTransactionsByUserIdAndAccountId(userId string, accountId string) ([]dto.Transaction, string) {
	var transactions []dto.Transaction
	db := config.DBConn()
	var query = "SELECT t.id_transaction, t.id_account, t.amount, a.bank, t.transaction_type, t.created_at " +
		"FROM dbbank.`transactions` t join dbbank.accounts a on t.id_account = a.id_account " +
		"WHERE a.id_user = " + userId
	if len(accountId) > 0 {
		query = query + " AND t.id_account = " + accountId
	}

	rows, err := db.Query(query)

	if err != nil {
		return transactions, "Query Db is failed"
	}

	for rows.Next() {
		transaction := dto.Transaction{}
		var id_transaction, account_id int
		var amount float64
		var bank, transaction_type string
		var created_at time.Time

		err = rows.Scan(&id_transaction, &account_id, &amount, &bank, &transaction_type, &created_at)
		if err != nil {
			return transactions, "No record in db"
		}

		transaction.IdTransaction = id_transaction
		transaction.IdAccount = account_id
		transaction.Amount = amount
		transaction.Bank = bank
		transaction.TransactionType = transaction_type
		transaction.CreatedAt = created_at
		transactions = append(transactions, transaction)
	}

	defer db.Close()
	return transactions, ""
}

func SaveTransaction(requestModel dto.RequestTransaction) string {

	db := config.DBConn()
	insTran, err := db.Prepare(
		"INSERT INTO transactions (id_account, amount, transaction_type) " +
			"VALUES(?, ?, ?)")
	if err != nil {
		//c.JSON(500, gin.H{
		//	"messages" : err,
		//})
		return "Insert is invalid"
	}
	insTran.Exec(requestModel.AccountId, requestModel.Amount, requestModel.TransactionType)
	defer db.Close()
	return "success"
}
