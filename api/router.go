package api

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"github.com/shinYeongHyeon/TodoList/db"
	"github.com/shinYeongHyeon/TodoList/todo"
	"net/http"
)

// TodoListAPI Hello world
func TodoListAPI() http.Handler {
	router := mux.NewRouter()

	router.Use(handlePanic)

	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)
	router.HandleFunc("/list", createTodoList).Methods(http.MethodPost)

	log.Info("Server Up Complete...")

	return router
}


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