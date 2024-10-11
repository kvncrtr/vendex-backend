package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := os.Getenv("CONNECTING_STRING")
	if connStr == "" {
		log.Fatal("DB_CONN_STRING not set in .env")
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Could not connect to database. Try again later.")
	}

	err = DB.Ping()
	if err != nil {
		panic("Could not ping database.")
	}

	log.Println("Successfully connected to the database!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
