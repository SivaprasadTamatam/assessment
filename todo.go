package main

import "fmt"

type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func DisplayToDo(todo ToDo) {
	fmt.Printf("ID: %d,	Title: %s,	Completed: %v\n", todo.ID, todo.Title, todo.Completed)
}
