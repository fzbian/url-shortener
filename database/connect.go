package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	env := godotenv.Load()
	if env != nil {
		log.Fatalf("err loading: %v", env)
	}

	var (
		db  *sql.DB
		err error
	)

	dbParams := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"))

	db, err = sql.Open("mysql", dbParams)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
