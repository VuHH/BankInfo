package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func DBConn() (db *sql.DB) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME_DATABASE")
	dbHost := os.Getenv("DB_HOST_DATABASE")
	dbPort := os.Getenv("DB_PORT_DATABASE")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())
		fmt.Println("err", err.Error())
	}
	return db
}
