package postgres

import (
	"log"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestDb(t *testing.T) {
	db := InitDB()
	err := db.Ping()
	check(err)
	row := db.QueryRow("select 1;")
	var res int
	check(row.Scan(&res))
	log.Print(res)
}
