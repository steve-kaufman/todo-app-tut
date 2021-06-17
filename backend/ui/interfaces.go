package ui

import "github.com/steve-kaufman/todo-app-tut/backend/entities"

type Service interface {
	GetAllTodos() ([]entities.Todo, error)
}
