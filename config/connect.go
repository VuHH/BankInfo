package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "P@ssword123++"
	dbName := "dbbank"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())
		fmt.Println("err", err.Error())
	}
	return db
}
