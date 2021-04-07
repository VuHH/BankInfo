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
	}
	return t

}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	//encodedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	//return  jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
	//	return []byte("SECRETKEY_123"), nil
	//})
	//
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return "", fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte("SECRETKEY_123"), nil
	})

}
