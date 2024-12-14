package app

import (
	"database/sql"
	"fmt"
	"net/http"

	"todo-api/pkg/config"
	"todo-api/pkg/db"
	"todo-api/pkg/server"

	"todo-api/internal/handlers/todo"
	repository "todo-api/internal/repositories"
	services "todo-api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type App struct {
	config *config.Config
	db     *sql.DB
}

func NewApp(config *config.Config) (*App, error) {
	// initialize db
	pgDb, err := db.NewPgDB(config.PGConfig)
	if err != nil {
		panic(err)
	}

	return &App{
		config: config,
		db:     pgDb,
	}, nil
}

func (a *App) Run() {
	// init newApp
	app, err := NewApp(a.config)
	if err != nil {
		panic(err)
	}

	// Initialize the TodoRepository
	todoRepo := repository.NewTodoRepository(app.db)

	// Initialize the TodoService
	todoService := services.NewTodoService(todoRepo)

	// Initialize the TodoHandler
	todoHandler := todo.NewTodoHandler(todoService)

	// create new router
	r := chi.NewRouter()

	// define cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// define routes
	server.GetRoutes(r, todoHandler)

	// start server
	port := 3000
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
