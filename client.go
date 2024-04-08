package main

// ToDoClient interface defines the method for Getting ToDo.
// We can extend this for different ways to get the ToDo like http,
// and also in any other ways to just not change the code code
type ToDoClient interface {
	GetTodo(id int) (ToDo, error)
}
