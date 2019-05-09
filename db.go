package spork

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
}

//GetDB - returns db instance
func GetDB() *sql.DB {
	return db
}
