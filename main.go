package main

import "fmt"

const ToDoBaseURL = "https://jsonplaceholder.typicode.com/todos/"
const MAX_COUNT = 20

func main() {
	var todos []ToDo
	client := HttpClient{}

	index := 1
	for {
		if len(todos) >= MAX_COUNT {
			break
		}
		todo, err := client.GetTodo(index)
		index++
		if err != nil {
			fmt.Println(err)
			break
		}
		if todo.ID%2 == 0 {
			todos = append(todos, todo)
		}
	}

	for _, todo := range todos {
		fmt.Printf("ID: %d,	Title: %s,	Completed: %v\n", todo.ID, todo.Title, todo.Completed)
	}
}
