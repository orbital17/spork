package spork

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
	db := GetDB()
	err := db.Ping()
	check(err)
	row := db.QueryRow("select 1;")
	var res int
	check(row.Scan(&res))
	log.Print(res)
}

func TestCreateUser(t *testing.T) {
	id, err := CreateUser(
		"olexiy.tkachenko+2@gmail.com",
		"testpassword",
		"o.tkachenkp",
	)
	check(err)
	if id <= 0 {
		t.Fail()
	}
	log.Printf("user id: %v", id)
}

func TestLoginFailed(t *testing.T) {
	token, err := Login(
		"olexiy.tkachenko@gmail.com",
		"incorrectpassword",
	)
	if err == nil || token != "" {
		t.Fail()
	}
}

func TestLogin(t *testing.T) {
	token, err := Login(
		"olexiy.tkachenko@gmail.com",
		"testpassword",
	)
	check(err)
	if len(token) == 0 {
		t.Fail()
	}
	log.Printf("user token: %v", token)
}

func TestNewToken(t *testing.T) {
	token, err := NewToken(User{
		id: 3,
	})
	check(err)
	t.Logf("user token: %v", token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjN9.hqg8qt7gaEfrxaPujYENQ8pSaO-stU_LW1tzRNVgGx8"
	id, err := ParseToken(token)
	check(err)
	t.Logf("user id: %v", id)
}
