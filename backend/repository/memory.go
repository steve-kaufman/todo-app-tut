package repository

import "github.com/steve-kaufman/todo-app-tut/backend/entities"

type Memory struct {
	todos []entities.Todo
}

func NewMemory(todos []entities.Todo) *Memory {
	return &Memory{
		todos: todos,
	}
}

func (repo Memory) GetAllTodos() ([]entities.Todo, error) {
	return repo.todos, nil
}
