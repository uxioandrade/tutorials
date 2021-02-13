package main

import (
	"log"

	"github.com/uxioandrade/go-sql-tutorial/model/todos"
)

func main() {
	firstTodo := todos.Todo{
		Description: "First todo",
		Priority:    1,
		Status:      "In Progress",
	}
	secondTodo := todos.Todo{
		Description: "Second todo",
		Priority:    3,
		Status:      "Done",
	}
	oldTodo := todos.Todo{
		ID: 1,
	}
	firstTodo.Save()
	secondTodo.Save()
	oldTodo.Get()
	log.Println(oldTodo)
}
