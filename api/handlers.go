package api

import (
	"github.com/shinYeongHyeon/TodoList/db"
	"github.com/shinYeongHyeon/TodoList/todo"
	"net/http"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()

	must(err)
	writeJSON(w, lists)
}

func createTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	todoList, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, todoList)
}