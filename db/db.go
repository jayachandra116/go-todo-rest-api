package db

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Open database 'api.db' with driver 'sqlite'
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Could not open database: " + err.Error())
	}
	// Verify connection to database
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Set maximum number of open connections to the database
	DB.SetMaxOpenConns(10)
	// Set maximum number of connections to the database in the idle connection pool
	DB.SetMaxIdleConns(5)
	fmt.Println("Connected to database successfully!")
	// Create database tables if they don't exist
	createTables()
}

func createTables() {
	createTodoTable()
	createUserTable()
}

func createTodoTable() {
	// Create todo table
	todoTableQueryString := `
	CREATE TABLE IF NOT EXISTS todos(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL
	)
	`
	_, err := DB.Exec(todoTableQueryString)
	if err != nil {
		panic("Could not create todo table: " + err.Error())
	}
}

func createUserTable() {
	// Create user table
	userTableQueryString := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(userTableQueryString)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}
}
