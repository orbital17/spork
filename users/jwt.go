package users

import (
	"strconv"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

var hs256 = jwt.NewHMAC(jwt.SHA256, []byte("jwtsecret"))

func NewToken(user User) (string, error) {
	now := time.Now()
	h := jwt.Header{}
	p := jwt.Payload{
		Subject:        strconv.Itoa(int(user.Id)),
		ExpirationTime: now.Add(24 * 30 * time.Hour).Unix(),
		IssuedAt:       now.Unix(),
	}
	token, err := jwt.Sign(h, p, hs256)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func ParseToken(tokenString string) (UserID, error) {
	raw, err := jwt.Parse([]byte(tokenString))
	if err != nil {
		return 0, err
	}
	if err = raw.Verify(hs256); err != nil {
		return 0, err
	}
	var (
		p jwt.Payload
	)
	_, err = raw.Decode(&p)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now, true)
	if err := p.Validate(iatValidator, expValidator); err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(p.Subject)
	if err != nil {
		return 0, err
	}
	return UserID(id), nil
}
