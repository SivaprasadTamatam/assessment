package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// httpClient is an implementation of ToDoClient interface using Http request
type HttpClient struct{}

// GetTodo gets a todo with a given id using HTTP GET request
func (hc *HttpClient) GetTodo(id int) (ToDo, error) {
	var result ToDo

	// Construct url
	url := fmt.Sprintf("%s%d", ToDoBaseURL, id)
	resp, err := http.Get(url)
	// Handle error
	if err != nil {
		return result, fmt.Errorf("error while getting todo for ID:%d :- %v", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("error while getting todo for ID:%d :with status %s", id, resp.Status)
	}

	// Parsing body response and construct ToDo struct
	if err := json.NewDecoder((resp.Body)).Decode(&result); err != nil {
		return result, fmt.Errorf("error while decoding response body in JSON format for ID: %d :- %v", id, err)
	}

	return result, nil
}
