package todos

import (
	"fmt"
	"log"

	"github.com/uxioandrade/go-sql-tutorial/datasources/postgres/todos_db"
)

func (todo *Todo) Get() error {
	stmt, err := todos_db.Client.Prepare("SELECT id, description, priority, status FROM todos WHERE id=$1;")
	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to prepare statement %s", err.Error()))
		log.Println(err)
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(todo.ID)

	if err := result.Scan(&todo.ID, &todo.Description, &todo.Priority, &todo.Status); err != nil {
		log.Println("Error when trying to get Todo by ID")
		return err
	}
	return nil
}

func (todo *Todo) Save() error {
	stmt, err := todos_db.Client.Prepare("INSERT INTO todos(description, priority, status) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		log.Println("Error when trying to prepare statement")
		log.Println(err)
		return err
	}
	defer stmt.Close()
	var lastInsertID int64
	insertErr := stmt.QueryRow(todo.Description, todo.Priority, todo.Status).Scan(&lastInsertID)
	if insertErr != nil {
		log.Println("Error when trying to save todo")
		return err
	}
	todo.ID = lastInsertID
	log.Println(fmt.Sprintf("Successfully inserted new todo with id %d", todo.ID))
	return nil
}
