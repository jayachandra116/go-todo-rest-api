package models

import (
	"example.com/todo/db"
)

type ToDo struct {
	ID      int64
	Content string
}

func (t *ToDo) Save() error {
	query := `
	INSERT INTO todos(content) VALUES (?)
	`
	// Create a prepared statement with the query string
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// Close the statement after this whole func is finished
	defer stmt.Close()
	// Execute the prepared statement with arguments from the todo item
	result, err := stmt.Exec(t.Content)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

func GetAllTodos() ([]ToDo, error) {
	query := `
	SELECT * FROM todos
	`
	// Execute the query
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
	query := `
	SELECT * FROM todos
	WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)
	var todo ToDo
	err := row.Scan(&todo.ID, &todo.Content)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *ToDo) Update() error {
	query := `
	UPDATE todos 
	SET content = ?
	WHERE ID = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(query, t.Content, t.ID)
	return err
}

func (t *ToDo) Delete() error {
	query := `
	DELETE FROM todos WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(query, t.ID)
	if err != nil {
		return err
	}
	return nil
}
