package repository

import (
	"bankinfo.com/config"
)

func UpdateNameAccount(nameBank string, accountId string) string {
	db := config.DBConn()
	upAccount, err := db.Prepare("UPDATE accounts SET name=? WHERE id_account= " + accountId)
	if err != nil {
		return "Update is invalid"
	}
	upAccount.Exec(nameBank)
	defer db.Close()
	return "success"
}

func UpdateStatusAccount(isActive bool, accountId string) string {
	db := config.DBConn()
	upAccount, err := db.Prepare("UPDATE accounts SET is_active=? WHERE id_account= " + accountId)
	if err != nil {
		return "Update is invalid"
	}
	upAccount.Exec(isActive)
	defer db.Close()
	return "success"
}
