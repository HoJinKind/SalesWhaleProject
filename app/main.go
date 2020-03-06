package main

import (
	"SalesWhaleProject/github.com/dgrijalva/jwt-go"
	_ "SalesWhaleProject/models"
	"fmt"
	. "strconv"
	"strings"
	"time"
)

func main() {
	//ls := []int {1,2,3}
	reg := []string {Itoa(1),Itoa(2),Itoa(3)}
	fmt.Println(strings.Join(reg[:], ","))
	st := strings.Join(reg[:], ",")
	arr := strings.Split(st, ",")
	fmt.Print(arr[0])

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": "bar",
		"nbf": time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("secret"))
	fmt.Println(tokenString)


	// test check

	tokenString1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImJhciIsIm5iZiI6MTYwMjMzMTIwMH0.w27KF5XT_c4PsXhCPvO7mwCEBs_mgkJt9XvOKanClOo"


	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString1, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		var hmacSampleSecret []byte
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	print(token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(123)
		fmt.Println(err)
	}
	hmacSecretreturn , valid := extractClaims(tokenString1)
	print(valid)
	if valid{
		print(hmacSecretreturn["nbf"])
	}
}

