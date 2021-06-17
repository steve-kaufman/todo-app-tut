package usecases

import "github.com/steve-kaufman/todo-app-tut/backend/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error) {
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal
	}
	return todos, nil
}
