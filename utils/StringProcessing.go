package utils

import (
	"fmt"
	"strings"
)

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