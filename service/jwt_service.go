package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

func GenerateToken(username string, isUser bool) string {

	claims := &authCustomClaims{
		username,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "vuhh",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("SECRETKEY_123"))
	if err != nil {
		panic(err)
		fmt.Println("err", err)
	}
	return t

}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return "", fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte("SECRETKEY_123"), nil
	})
}
