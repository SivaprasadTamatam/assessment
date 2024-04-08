package main

// ToDoClient interface defines the method for Getting ToDo.
type ToDoClient interface {
	GetTodo(id int) (ToDo, error)
}
