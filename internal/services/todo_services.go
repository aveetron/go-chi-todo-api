package service

import (
	"fmt"
	"log"
	"todo-api/internal/models"
	repository "todo-api/internal/repositories"
)

type TodoService struct {
	TodoRepository repository.TodoRepository
}

// NewTodoService creates and initializes a TodoService with a TodoRepository.
func NewTodoService(repo repository.TodoRepository) *TodoService {
	if repo == nil {
		panic("TodoRepository is not initialized")
	}
	return &TodoService{
		TodoRepository: repo,
	}
}

func (ts *TodoService) GetAllTodos() ([]*models.TodoSchema, error) {
	if ts.TodoRepository == nil {
		log.Println("Error: TodoRepository is nil")
		return nil, fmt.Errorf("TodoRepository is nil")
	}

	todos, err := ts.TodoRepository.GetAllTodos()
	if err != nil {
		log.Printf("Error fetching todos: %v", err)
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}

	return todos, nil
}

func (ts *TodoService) CreateTodo(todo *models.TodoSchema) error {
	return ts.TodoRepository.CreateTodo(todo)
}

func (ts *TodoService) GetTodoByID(id int) (*models.TodoSchema, error) {
	return ts.TodoRepository.GetTodoByID(id)
}

func (ts *TodoService) UpdateTodoByID(id int) (*models.TodoSchema, error) {
	return ts.TodoRepository.UpdateTodoByID(id)
}

func (ts *TodoService) DeleteTodoByID(id int) error {
	return ts.TodoRepository.DeleteTodoByID(id)
}
