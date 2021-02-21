package api

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
)

// TodoListAPI Hello world
func TodoListAPI() http.Handler {
	router := mux.NewRouter()

	router.Use(handlePanic)

	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)
	router.HandleFunc("/list/{list_id}", getTodoList).Methods(http.MethodGet)

	router.HandleFunc("/list", createTodoList).Methods(http.MethodPost)
	router.HandleFunc("/list/{list_id}/item", createTodoItem).Methods(http.MethodPost)

	router.HandleFunc("/list/{list_id}", renameTodoList).Methods(http.MethodPut)
	router.HandleFunc("/list/{list_id}/item/{item_id}", modifyTodoItem).Methods(http.MethodPut)

	router.HandleFunc("/list/{list_id}", deleteTodoList).Methods(http.MethodDelete)

	log.Info("Server Up Complete...")

	return router
}
