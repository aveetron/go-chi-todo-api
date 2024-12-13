package models

type TodoSchema struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	IsDone      bool   `db:"is_done" json:"status" default:"false"`
}
