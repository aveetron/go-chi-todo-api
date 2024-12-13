package db

import (
	"database/sql"
	"fmt"

	"todo-api/pkg/config"

	_ "github.com/lib/pq"
)

func NewPgDB(cfg config.PGConfig) (*sql.DB, error) {
	connectionStrig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBNAME,
	)

	db, err := sql.Open("postgres", connectionStrig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
