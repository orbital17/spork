package auth

import (
	"log"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func TestToken(t *testing.T) {
	id := int64(345)
	token, err := NewToken(Auth{
		UserID: id,
	})
	check(err)
	// t.Logf("user token: %v", token)
	auth, err := ParseToken(token)
	check(err)
	if id != auth.UserID {
		t.Fail()
	}
	// t.Logf("user id: %v", id)
}
