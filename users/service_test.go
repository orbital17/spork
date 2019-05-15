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

func TestToken(t *testing.T) {
	token, err := NewToken(User{
		id: 345,
	})
	check(err)
	t.Logf("user token: %v", token)
	id, err := ParseToken(token)
	check(err)
	t.Logf("user id: %v", id)
}
