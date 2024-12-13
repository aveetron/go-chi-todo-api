package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-api/internal/models"
	service "todo-api/internal/services"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	todoService *service.TodoService
}

func NewTodoHandler(todoService *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (th *TodoHandler) HandleGetAllTodos(w http.ResponseWriter, r *http.Request) {
	// Add nil check for todoService
	if th.todoService == nil {
		http.Error(w, "Todo service not initialized", http.StatusInternalServerError)
		return
	}

	todos, err := th.todoService.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to retrieve todos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (th *TodoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	if th.todoService == nil {
		http.Error(w, "Todo service not initialized", http.StatusInternalServerError)
		return
	}

	todo := &models.TodoSchema{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Optional: Add basic validation
	if todo.Title == "" {
		http.Error(w, "Todo title is required", http.StatusBadRequest)
		return
	}

	if err := th.todoService.CreateTodo(todo); err != nil {
		http.Error(w, "Failed to create todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (th *TodoHandler) HandleGetTodoByID(w http.ResponseWriter, r *http.Request) {
	if th.todoService == nil {
		http.Error(w, "Todo service not initialized", http.StatusInternalServerError)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid todo ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := th.todoService.GetTodoByID(id)
	if err != nil {
		fmt.Errorf("Failed to update todo: %w", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (th *TodoHandler) HandleUpdateTodoByID(w http.ResponseWriter, r *http.Request) {
	if th.todoService == nil {
		http.Error(w, "Todo service not initialized", http.StatusInternalServerError)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid todo ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := th.todoService.UpdateTodoByID(id)
	if err != nil {
		fmt.Errorf("Failed to update todo: %w", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (th *TodoHandler) HandleDeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	if th.todoService == nil {
		http.Error(w, "Todo service not initialized", http.StatusInternalServerError)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid todo ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := th.todoService.DeleteTodoByID(id); err != nil {
		fmt.Errorf("Failed to update todo: %w", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
