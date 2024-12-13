package todo

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"status" default:"false"`
}

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	IsDone      bool   `json:"status"`
}
