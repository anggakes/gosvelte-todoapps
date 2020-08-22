package models

type Todo struct {
	ID          int `json:"id" `
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoCreateCommand struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoUpdateCommand struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoGetQuery struct {
	ID string `json:"id"`
}

type TodoListQuery struct {
	Size int `json:"size"`
	Page int `json:"page"`
}

type TodoListResults struct {
	CurrentPage int    `json:"current_page"`
	Results     []Todo `json:"results"`
}
