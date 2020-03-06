package models

import (
	_ "SalesWhaleProject/github.com/go-sqlite3"
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	db = openConnection()
	query := `
	CREATE TABLE IF NOT EXISTS boards (id TEXT, token TEXT, end_epoch INTEGER,duration INTEGER, board_array TEXT, points INTEGER, node_array TEXT );
	CREATE TABLE IF NOT EXISTS node (id TEXT, value TEXT, adjacent_array TEXT);
	`
	db.Exec(query)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func openConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		log.Panic(err)
	}
	return db
}