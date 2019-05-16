package users

import (
	"context"
	"strconv"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

var hs256 = jwt.NewHMAC(jwt.SHA256, []byte("jwtsecret"))

type Auth struct {
	UserID UserID
}

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

func ParseToken(tokenString string) (auth *Auth, err error) {
	raw, err := jwt.Parse([]byte(tokenString))
	if err != nil {
		return
	}
	if err = raw.Verify(hs256); err != nil {
		return
	}
	var (
		p jwt.Payload
	)
	_, err = raw.Decode(&p)
	if err != nil {
		return
	}
	now := time.Now()
	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now, true)
	if err = p.Validate(iatValidator, expValidator); err != nil {
		return
	}
	id, err := strconv.Atoi(p.Subject)
	if err != nil {
		return
	}
	return &Auth{UserID(id)}, nil
}

type contextKey int

var authContextKey contextKey

func NewContext(ctx context.Context, auth *Auth) context.Context {
	return context.WithValue(ctx, authContextKey, auth)
}

func FromContext(ctx context.Context) (*Auth, bool) {
	auth, ok := ctx.Value(authContextKey).(*Auth)
	return auth, ok
}
