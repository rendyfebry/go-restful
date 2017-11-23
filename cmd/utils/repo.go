package utils

import (
	"fmt"

	"github.com/rendyfebry/go-restful/cmd/models"
)

var currentID int

var TodoList models.Todos

// Create seed data
func init() {
	RepoCreateTodo(models.Todo{Name: "Write presentation"})
	RepoCreateTodo(models.Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) models.Todo {
	for _, t := range TodoList {
		if t.Id == id {
			return t
		}
	}

	// return empty Todo if not found
	return models.Todo{}
}

func RepoCreateTodo(t models.Todo) models.Todo {
	currentID++
	t.Id = currentID
	TodoList = append(TodoList, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range TodoList {
		if t.Id == id {
			TodoList = append(TodoList[:i], TodoList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
