package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const dbConnStr = "postgres://postgres:postgres@localhost:5434/todo_db?sslmode=disable"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run migrations/main.go [up|down]")
		os.Exit(1)
	}

	command := os.Args[1]

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	switch command {
	case "up":
		err = Up(db)
	case "down":
		err = Down(db)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		fmt.Printf("Migration %s completed successfully.\n", command)
	}
}

func Up(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		is_done BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create todos table: %w", err)
	}
	return nil
}

func Down(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS todos;`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to drop todos table: %w", err)
	}
	return nil
}
