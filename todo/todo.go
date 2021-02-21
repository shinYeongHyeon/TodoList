package todo

// List a todo list
type List struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

// ListWithItems a todo list with its items
type ListWithItems struct {
	List
	Items []Item `json:"items"`
}

// Item a todo list item
type Item struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
