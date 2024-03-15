package models

import (
	"example.com/todo/db"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password) VALUES (?, ?)
	`
	// Create a prepared statement with the query string
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// Close the statement after this whole func is finished
	defer stmt.Close()
	// Execute the prepared statement with arguments from the todo item
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (u *User) Update() error {
	query := `
	UPDATE users
	SET email = ?, password = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Email, u.Password, u.ID)
	return err
}

func (u *User) Delete() error {
	query := `
	DELETE from users
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID)
	return err
}

func GetUserById(id int64) (*User, error) {
	query := `
		SELECT * FROM users
		WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT * FROM users
		WHERE email = ?
	`
	row := db.DB.QueryRow(query, email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ChangePassword(id int64, newPassword string) error {
	query := `
		UPDATE users SET password =? WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newPassword, id)
	return err
}
