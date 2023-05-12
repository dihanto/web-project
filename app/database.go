package app

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func GetConnection() *sql.DB {
	if Db == nil {
		var err error
		Db, err = sql.Open("mysql", "root@tcp(localhost:3306)/web")
		if err != nil {
			panic(err)
		}
	}
	return Db
}
