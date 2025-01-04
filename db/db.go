package db

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error

	DB, err = sql.Open("sqlite3", "api.bd")

	if err != nil {
		panic("error opening database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()

}



func createTable() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		is_admin BOOLEAN DEFAULT FALSE
	)
	`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("could not create user table")
	}


	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic("could not create event table: " + err.Error())
		
	}

	createRegisterEventTable := `
	CREATE TABLE IF NOT EXISTS registration (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		event_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (event_id) REFERENCES events(id)
		)`

	_, err = DB.Exec(createRegisterEventTable)

	if err != nil {
		panic("could not create registration table: " + err.Error())
	}

}