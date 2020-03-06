package models

import (
	"fmt"
)

type Person struct {
	Id        string
	Firstname string
	Lastname  string
}

func AllBooks() ([]*Person, error) {
	rows, err := db.Query("SELECT Id, Firstname, Lastname FROM people")
	if err != nil {
		return nil, err
	}

	bks := make([]*Person, 0)

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM people")
	row.Scan(&count)
	fmt.Print(count)

	for rows.Next() {

		bk := new(Person)
		rows.Scan(&bk.Id, &bk.Firstname, &bk.Lastname)
		fmt.Println(bk.Id + ": " + bk.Firstname + " " + bk.Lastname)

		err := rows.Scan(&bk.Id, &bk.Firstname, &bk.Lastname)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}


