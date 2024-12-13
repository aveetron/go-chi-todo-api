package server

import (
	"todo-api/internal/handlers/health"
	"todo-api/internal/handlers/todo"

	"github.com/go-chi/chi/v5"
)

func GetRoutes(r *chi.Mux, todoHandler *todo.TodoHandler) {

	r.Get("/healthCheck", health.HandleHealthCheck)
	r.Get("/todos", todoHandler.HandleGetAllTodos)
	r.Post("/todos/", todoHandler.HandleCreateTodo)
	r.Get("/todos/{id}", todoHandler.HandleGetTodoByID)
	r.Put("/todos/{id}", todoHandler.HandleUpdateTodoByID)
	r.Delete("/todos/{id}", todoHandler.HandleDeleteTodoByID)
}
