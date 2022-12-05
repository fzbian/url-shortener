package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	var (
		db  *sql.DB
		err error
	)

	db, err = sql.Open("mysql", "root:@/url_shortener")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
