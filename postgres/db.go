package postgres

import (
	"database/sql"
	"log"
	"spork/config"

	_ "github.com/lib/pq"
)

func Init(config *config.Config) *sql.DB {
	db, err := sql.Open("postgres", config.PostgresConnStr)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	//TODO: test query
	return db
}
