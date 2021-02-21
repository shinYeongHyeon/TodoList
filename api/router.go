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
	router.HandleFunc("/list", createTodoList).Methods(http.MethodPost)

	log.Info("Server Up Complete...")

	return router
}
