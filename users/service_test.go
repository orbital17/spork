package users

import (
	"log"
	"spork/postgres"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var service Service

func init() {
	db := postgres.InitDB()
	store := NewStore(db)
	service = NewService(store)
}

func TestCreateUser(t *testing.T) {
	id, err := service.CreateUser(
		"olexiy.tkachenko+3@gmail.com",
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
	token, err := service.Login(
		"olexiy.tkachenko@gmail.com",
		"incorrectpassword",
	)
	if err == nil || token != "" {
		t.Fail()
	}
}

func TestLogin(t *testing.T) {
	token, err := service.Login(
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
