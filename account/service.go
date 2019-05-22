package account

import (
	"errors"
	"regexp"
	"spork/auth"
	"spork/users"

	"golang.org/x/crypto/bcrypt"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var ErrBadEmailFormat = errors.New("invalid email format")

type Service struct {
	store *users.Store
}

func NewService(store *users.Store) *Service {
	return &Service{store}
}

func (service *Service) Create(
	email string,
	password string,
	name string,
) (
	id int64,
	err error,
) {
	if !emailRegexp.MatchString(email) {
		return 0, ErrBadEmailFormat
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	return service.store.AddUser(users.User{
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
	token, err = auth.NewToken(auth.Auth{user.Id})
	return
}
