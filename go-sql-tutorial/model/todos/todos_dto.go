package todos

type Todo struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
}
