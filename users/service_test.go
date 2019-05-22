package users

import (
	"log"
	"spork/config"
	"spork/postgres"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var service *Service

func init() {
	config := config.Default()
	db := postgres.Init(config)
	store := NewStore(db)
	service = NewService(store)
}

func TestCreateUser(t *testing.T) {
	id, err := service.CreateUser(
		"olexiy.tkachenko+3@gmail.com",
		"testpassword",
		"o.tkachenkp",
	)
	if err != nil && err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
		t.SkipNow()
	}
	check(err)
	if id <= 0 {
		t.Fail()
	}
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
	// log.Printf("user token: %v", token)
}
