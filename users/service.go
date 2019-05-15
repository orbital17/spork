package users

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var ErrBadEmailFormat = errors.New("invalid email format")

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store}
}

func (service *Service) CreateUser(
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
	return service.store.AddUser(User{
		Email:        email,
		PasswordHash: string(hash),
		Name:         name,
	})
}

func (service *Service) Login(
	email string,
	password string,
) (
	token string,
	err error,
) {
	user, err := service.store.UserByEmail(email)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
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
