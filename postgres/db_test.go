package postgres

import (
	"database/sql"
	"log"
	"spork/config"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDb() *sql.DB {
	config := config.Default()
	return Init(config)
}

func TestDb(t *testing.T) {
	db := getDb()
	err := db.Ping()
	check(err)
	row := db.QueryRow("select 1;")
	var res int
	check(row.Scan(&res))
	log.Print(res)
}
