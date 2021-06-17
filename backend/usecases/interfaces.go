package usecases

import "github.com/steve-kaufman/todo-app-tut/backend/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
