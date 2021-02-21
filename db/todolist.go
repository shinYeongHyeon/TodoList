package db

import (
	"github.com/shinYeongHyeon/TodoList/todo"
)

// GetTodoLists returns all todo lists
func GetTodoLists() ([]todo.List, error) {
	rows, err := db.Query(`SELECT id, name FROM todo_list`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := []todo.List{}
	for rows.Next() {
		var list todo.List
		if err := rows.Scan(&list.ID, &list.Name); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func CreateTodoList(name string) (list todo.List, err error) {
	list.Name = name
	err = db.QueryRow(`
		INSERT INTO todo_list (name) VALUES ($1) RETURNING id
	`, name).Scan(&list.ID)
	return
}