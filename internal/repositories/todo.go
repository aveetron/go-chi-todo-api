package repository

import (
	"database/sql"
	"time"
	"todo-api/internal/models"
)

type TodoRepository interface {
	GetAllTodos() ([]*models.TodoSchema, error)
	CreateTodo(*models.TodoSchema) error
	GetTodoByID(id int) (*models.TodoSchema, error)
	UpdateTodoByID(id int, todo *models.TodoSchema) (*models.TodoSchema, error)
	DeleteTodoByID(id int) error
}

type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &TodoRepositoryImpl{db: db}
}

func (tr *TodoRepositoryImpl) GetAllTodos() ([]*models.TodoSchema, error) {
	query := "SELECT * FROM todos"

	rows, err := tr.db.Query(query)
	if err != nil {
		return nil, err
	}

	var todos []*models.TodoSchema
	for rows.Next() {
		todo := &models.TodoSchema{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone, &todo.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	defer rows.Close()
	return todos, nil
}

func (tr *TodoRepositoryImpl) CreateTodo(todo *models.TodoSchema) error {
	query := "INSERT INTO todos (title, description, is_done, created_at) VALUES ($1, $2, $3, $4)"
	_, err := tr.db.Exec(query, todo.Title, todo.Description, todo.IsDone, time.Now())
	return err
}

func (tr *TodoRepositoryImpl) GetTodoByID(id int) (*models.TodoSchema, error) {
	query := "SELECT * FROM todos WHERE id = $1"
	todo := &models.TodoSchema{}
	if err := tr.db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone, &todo.CreatedAt); err != nil {
		return nil, err
	}
	return todo, nil
}

func (tr *TodoRepositoryImpl) UpdateTodoByID(id int, todo *models.TodoSchema) (*models.TodoSchema, error) {
	query := "UPDATE todos SET title = $1, description = $2, is_done = $3 WHERE id = $4 RETURNING id, title, description, is_done, created_at"

	err := tr.db.QueryRow(query, todo.Title, todo.Description, todo.IsDone, id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (tr *TodoRepositoryImpl) DeleteTodoByID(id int) error {
	query := "DELETE FROM todos WHERE id = $1"
	_, err := tr.db.Exec(query, id)
	return err
}
