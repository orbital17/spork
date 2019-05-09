package users

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var secret []byte = []byte("jwtsecret")

func NewToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.id,
	})
	return token.SignedString(secret)
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if m, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || m != jwt.SigningMethodHS256 {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return secret, nil
}

func ParseToken(tokenString string) (UserID, error) {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return UserID(claims["userId"].(int64)), nil
	} else {
		return 0, err
	}

}
