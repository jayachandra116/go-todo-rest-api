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
	fmt.Println("Opening database 'api.db' ...")
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Could not open database: " + err.Error())
	}
	// Verify connection to database
	fmt.Println("Verifying connection to database...")
	err = DB.Ping()
	if err != nil {
		panic("Could not verify connection to  database: " + err.Error())
	}
	// Set maximum number of open connections to the database
	fmt.Println("Setting maximum number of open connections to 10...")
	DB.SetMaxOpenConns(10)
	// Set maximum number of connections to the database in the idle connection pool
	fmt.Println("Setting maximum number of connections to 5...")
	DB.SetMaxIdleConns(5)
	fmt.Println("Connected to database successfully!")
	// Create database tables if they don't exist
	createTables()
}

func createTables() {
	fmt.Println("Creating tables...")
	createUserTable()
	createTodoTable()
}

func createTodoTable() {
	fmt.Println("Creating 'todos' table...")
	// Create todo table
	todoTableQueryString := `
	CREATE TABLE IF NOT EXISTS todos(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		userId INTEGER,
		FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
	)
	`
	_, err := DB.Exec(todoTableQueryString)
	if err != nil {
		panic("Could not create todo table: " + err.Error())
	}
}

func createUserTable() {
	fmt.Println("Creating 'users' table...")
	// Create user table
	userTableQueryString := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(userTableQueryString)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}
}
