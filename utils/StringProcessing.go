package utils

import (
	"fmt"
	"strings"
)

var DefaultBoard = "T, A, P, *, E, A, K, S, O, B, R, S, S, *, X, D"

func SplitString(str string) []string {
	fmt.Println(strings.Split(str, ", ")[15])
	return strings.Split(str, ", ")
}

func StringConcat(arr []string) string {
	return strings.Join(arr[:], ", ")
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetFirstParam(path string) (ps string) {

	// ignore first '/' and when it hits the second '/'
	// get whatever is after it as a parameter
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			ps = path[i+1:]
		}
	}
	return
}
