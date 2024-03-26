package models

import (
	"errors"
	"fmt"

	"example.com/todo/db"
	"example.com/todo/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Saves a new user to the database
func (u *User) Save() error {
	fmt.Println("Saving user ...")
	query := `
	INSERT INTO users(email, password) VALUES (?, ?)
	`
	// fmt.Println("Query to execute: ", query)
	// Create a prepared statement with the query string
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// Close the statement after this whole func is finished
	defer stmt.Close()
	fmt.Println("Executing query ...")
	// Execute the prepared statement with arguments from the todo item
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("Saved user id: ", id)
	u.ID = id
	return nil
}

func (u *User) Update() error {
	fmt.Println("Updating user...")
	query := `
	UPDATE users
	SET
	email = ?
	WHERE id = ?
	`
	fmt.Println("Query: ", query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println("Executing query ...")
	_, err = stmt.Exec(u.Email, u.ID)
	return err
}

func (u *User) Delete() error {
	fmt.Println("Deleting user with id", u.ID)
	query := `
	DELETE from users
	WHERE id = ?
	`
	fmt.Println("Query: ", query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println("Executing query ...")
	_, err = stmt.Exec(u.ID)
	return err
}

// Validates the user's credentials with email and password
func ValidateCredentials(email string, password string) error {
	fmt.Println("Validating user credentials ...")
	query := `
	SELECT password from users
	WHERE email = ?
	`
	// fmt.Println("Query:", query)
	fmt.Println("Executing query ...")
	row := db.DB.QueryRow(query, email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}
	fmt.Println("Checking password hash ...")
	passwordIsValid := utils.CheckPasswordHash(password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials Invalid")
	}
	return nil
}

func GetUserById(id int64) (*User, error) {
	fmt.Println("Getting user with id ", id)
	query := `
		SELECT id,email
		FROM users
		WHERE id = ?
	`
	// fmt.Println("Query to execute: ", query)
	fmt.Println("Executing query ...")
	row := db.DB.QueryRow(query, id)
	var user User
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Retrieve user by email
func GetUserByEmail(email string) (*User, error) {
	fmt.Println("Getting user by email with email: ", email)
	query := `
		SELECT id, email FROM users
		WHERE email = ?
	`
	// fmt.Println("Query: ", query)
	fmt.Println("Executing query ...")
	row := db.DB.QueryRow(query, email)
	var user User
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ChangeUserPassword(id int64, newPassword string) error {
	fmt.Println("Changing password of user with id: ", id)
	query := `
		UPDATE users SET password =? WHERE id = ?
	`
	fmt.Println("Query to execute: ", query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println("Executing query ...")
	_, err = stmt.Exec(newPassword, id)
	return err
}
