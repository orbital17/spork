package rest

import (
	"testing"
)

var client = NewClient()

func TestLogin(t *testing.T) {
	if !client.Login() {
		t.Fail()
	}
}
