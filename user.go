package spork

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type UserID int64

type User struct {
	id           UserID
	email        string
	name         string
	passwordHash string
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var ErrBadEmailFormat = errors.New("invalid email format")

func CreateUser(
	email string,
	password string,
	name string,
) (
	id UserID,
	err error,
) {
	if !emailRegexp.MatchString(email) {
		return 0, ErrBadEmailFormat
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	return newUser(User{
		email:        email,
		passwordHash: string(hash),
		name:         name,
	})
}

func Login(
	email string,
	password string,
) (
	token string,
	err error,
) {
	user, err := userByEmail(email)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.passwordHash), []byte(password))
	if err != nil {
		return
	}
	token, err = NewToken(user)
	return
}

func Auth(token string) (
	id UserID,
	err error,
) {
	return
}
