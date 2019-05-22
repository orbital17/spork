package rest

import (
	"testing"
)

var client = NewClient()

var creds = Credentials{
	Email:    "olexiy.tkachenko+3@gmail.com",
	Password: "testpassword",
}

func TestLogin(t *testing.T) {
	if !client.Login(creds) {
		t.Fail()
	}
}
