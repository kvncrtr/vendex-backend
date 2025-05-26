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

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Could not connect to database. Try again later.")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Could not ping database %v", err)
	}

	log.Println("Successfully connected to the database!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
