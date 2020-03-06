package main

import (
	"SalesWhaleProject/algo"
	_ "SalesWhaleProject/github.com/go-sqlite3"
	"SalesWhaleProject/models"
	"fmt"
)
func main() {
	fmt.Println(models.UpCheck(5,5))
	fmt.Println(models.RightDownCheck(0,5))
	fmt.Println(models.UpCheck(1,5))
	fmt.Println(models.UpCheck(1,5))
	str := "1234"
	fmt.Println(string(str[2]))
	//arr := []int{1,2,3}
	fmt.Print(algo.BoggleSolver( "A, C, E, D, D, D, D, D, D, D, D, D, D, D, D, D", "D"))
}

