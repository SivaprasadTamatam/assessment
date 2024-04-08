package main

import "fmt"

var ToDoBaseURL = "https://jsonplaceholder.typicode.com/todos/"

const MAX_COUNT = 20

func main() {

	client := HttpClient{}
	index := 1
	count := 0
	for {
		// breaking loop once we found even numbered TODO's of count 20
		if count >= MAX_COUNT {
			break
		}
		// getting TODO through http client
		todo, err := client.GetTodo(index)
		index++
		if err != nil {
			fmt.Println(err)
			break
		}
		// Display even numbered ToDOs
		if todo.ID%2 == 0 {
			count++
			DisplayToDo(todo)
		}
	}
}
