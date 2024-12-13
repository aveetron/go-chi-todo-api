package app

import (
	"fmt"
	"log"
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
}

func NewApp(config *config.Config) *App {
	// initialize db
	_, err := db.NewPgDB(config.PGConfig)
	if err != nil {
		panic(err)
	}

	return &App{
		config: config,
	}
}

func (a *App) Run() {
	// Load the configuration (from a config file or environment variables)
	cfg := config.PGConfig{
		Host:     "localhost",
		Port:     "5434",
		User:     "postgres",
		Password: "postgres",
		DBNAME:   "todo_db",
	}

	// Initialize the PostgreSQL database
	db, err := db.NewPgDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the TodoRepository
	todoRepo := repository.NewTodoRepository(db)

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
