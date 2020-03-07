package utils

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func TokenGenerator(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"nbf": time.Now().Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Panic(err)
	}
	return tokenString
}
