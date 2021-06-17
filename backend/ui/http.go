package ui

import (
	"encoding/json"
	"net/http"
)

type HTTPServer struct {
	service Service
}

func NewHTTP(service Service) *HTTPServer {
	return &HTTPServer{
		service: service,
	}
}

func (server HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, _ := server.service.GetAllTodos()
	json.NewEncoder(w).Encode(todos)
}
