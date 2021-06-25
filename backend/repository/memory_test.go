package repository_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/steve-kaufman/todo-app-tut/backend/entities"
	"github.com/steve-kaufman/todo-app-tut/backend/repository"
)

var tests = [][]entities.Todo{
	{},
	{
		{
			Title:       "todo 1",
			Description: "description",
			IsCompleted: true,
		},
	},
	{
		{
			Title:       "todo 1",
			Description: "description of todo 1",
			IsCompleted: true,
		},
		{
			Title:       "todo 2",
			Description: "description of todo 2",
			IsCompleted: false,
		},
		{
			Title:       "todo 3",
			Description: "description of todo 3",
			IsCompleted: true,
		},
	},
}

func TestMemory(t *testing.T) {
	for i, inputTodos := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			expect := expectate.Expect(t)

			repo := repository.NewMemory(inputTodos)

			todos, err := repo.GetAllTodos()

			expect(err).ToBe(nil)
			expect(todos).ToEqual(inputTodos)
		})
	}
}
