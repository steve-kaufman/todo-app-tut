package ui_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/steve-kaufman/todo-app-tut/backend/entities"
	"github.com/steve-kaufman/todo-app-tut/backend/ui"
)

type MockService struct {
	todos []entities.Todo
	err   error
}

func (s MockService) GetAllTodos() ([]entities.Todo, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.todos, nil
}

func TestHTTP(t *testing.T) {
	var expect expectate.ExpectorFunc
	var server *ui.HTTPServer
	var w *httptest.ResponseRecorder
	var r *http.Request

	setup := func(t *testing.T, service *MockService) {
		expect = expectate.Expect(t)
		server = ui.NewHTTP(service)
		w = httptest.NewRecorder()
		r = httptest.NewRequest(http.MethodGet, "http://mywebsite.com/", nil)
	}

	t.Run("Returns todos from root endpoint", func(t *testing.T) {
		setup(t, &MockService{todos: []entities.Todo{
			{
				Title:       "mock todo",
				Description: "description of mock todo",
				IsCompleted: false,
			},
		}})

		server.ServeHTTP(w, r)

		var body []entities.Todo
		json.NewDecoder(w.Result().Body).Decode(&body)

		expect(body).ToEqual([]entities.Todo{
			{
				Title:       "mock todo",
				Description: "description of mock todo",
				IsCompleted: false,
			},
		})
	})
}
