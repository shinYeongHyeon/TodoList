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

// GetTodoList returns a specific todo list with its items
func GetTodoList(todoListID int) (todo.ListWithItems, error) {
	var list todo.ListWithItems

	rows, err := db.Query(`
	SELECT tl.id, tl.name, ti.id, ti.text, ti.done
	FROM
	    todo_list tl LEFT JOIN
		todo_item ti
		ON tl.id = ti.todo_list_id
	WHERE tl.id = $1`, todoListID)
	if err != nil {
		return list, err
	}
	defer rows.Close()

	list.Items = []todo.Item{}
	gotTodoList := false
	for rows.Next() {
		var (
			itemID   *int
			itemText *string
			itemDone *bool
		)

		if err := rows.Scan(
			&list.ID,
			&list.Name,
			&itemID,
			&itemText,
			&itemDone); err != nil {
			return list, err
		}
		gotTodoList = true

		if itemID != nil && itemText != nil && itemDone != nil {
			list.Items = append(list.Items, todo.Item {
				ID:   *itemID,
				Text: *itemText,
				Done: *itemDone,
			})
		}
	}

	if !gotTodoList {
		return list, ErrNotFound
	}

	return list, nil
}

// CreateTodoList creates a new todo list
func CreateTodoList(name string) (list todo.List, err error) {
	list.Name = name
	err = db.QueryRow(`
		INSERT INTO todo_list (name) VALUES ($1) RETURNING id
	`, name).Scan(&list.ID)
	return
}

// RenameTodoList renames a todo list
func RenameTodoList(id int, newName string) error {
	res, err := db.Exec(`
	UPDATE todo_list SET name = $1 WHERE id = $2`, newName, id)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

// DeleteTodoList deletes a todo list
func DeleteTodoList(id int) error {
	res, err := db.Exec(`
	DELETE FROM todo_list WHERE id = $1`, id)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

// CreateTodoItem creates a new todo item
func CreateTodoItem(listID int, text string, done bool) (item todo.Item, err error) {
	item.Text = text
	item.Done = done
	err = db.QueryRow(`
	INSERT INTO todo_item
	    (todo_list_id, text, done)
	VALUES ($1, $2, $3) RETURNING id`,
	listID,
	text,
	done).Scan(&item.ID)
	return
}

// ModifyTodoItem modifies a todo item
func ModifyTodoItem(listID, itemID int, text string, done bool) error {
	res, err := db.Exec(`
	UPDATE todo_item
	SET text = $1, done = $2
	WHERE id = $3 AND todo_list_id = $4`,
	text,
	done,
	itemID,
	listID)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}