package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "db6"
	dbUser := "root"
	dbPass := "root"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	//db, err = sql.Open("mysql", "root:root@/db6")
	if err != nil {
		log.Fatal(err)
	}
	return
}
