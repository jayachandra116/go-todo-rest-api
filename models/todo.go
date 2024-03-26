package models

import (
	"fmt"

	"example.com/todo/db"
)

type ToDo struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	UserID  int64  `json:"userId"`
}

// Saves a new ToDo item to the database returns errors if any or nil
func (t *ToDo) Save() error {
	fmt.Println("Saving user to database ...")
	query := `INSERT INTO todos(content,userId) VALUES (?,?)`
	// Create a prepared statement with the query string
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// fmt.Println("Query Statement: ", stmt)
	// Close the statement after this whole func is finished
	defer stmt.Close()
	// Execute the prepared statement with arguments from the todo item
	fmt.Println("Executing prepared statement ...")
	result, err := stmt.Exec(t.Content, t.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("Last inserted id: ", id)
	t.ID = id
	return nil
}

func GetAllTodos() ([]ToDo, error) {
	fmt.Println("Getting all to-do items ...")
	query := `SELECT id, content, userId FROM todos`
	// Execute the query
	// fmt.Println("Query to execute:", query)
	fmt.Println("Executing query ...")
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []ToDo
	for rows.Next() {
		var todo ToDo
		err := rows.Scan(&todo.ID, &todo.Content)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodoById(id int64) (*ToDo, error) {
	fmt.Println("Getting todo by id with id: ", id)
	query := `
	SELECT id, content, userId FROM todos
	WHERE id = ?
	`
	// fmt.Println("Query to execute: ", query)
	fmt.Println("Executing query: ", query)
	row := db.DB.QueryRow(query, id)
	var todo ToDo
	err := row.Scan(&todo.ID, &todo.Content, &todo.UserID)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *ToDo) Update() error {
	fmt.Println("Updating the todo ...")
	query := `
	UPDATE todos 
	SET content = ?
	WHERE ID = ?
	`
	// fmt.Println("Query to execute: ", query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println("Executing query ...")
	_, err = stmt.Exec(query, t.Content, t.ID)
	return err
}

func (t *ToDo) Delete() error {
	fmt.Println("Deleting todo item ...")
	query := `
	DELETE FROM todos WHERE id = ?
	`
	// fmt.Println("Query to execute: ", query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println("Executing query ...")
	_, err = stmt.Exec(query, t.ID)
	if err != nil {
		return err
	}
	return nil
}
