package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"github.com/shinYeongHyeon/TodoList/db"
	"net/http"
)

// TodoListAPI Hello world
func TodoListAPI() http.Handler {
	router := mux.NewRouter()

	router.Use(handlePanic)

	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)

	log.Info("Server Up Complete...")

	return router
}


func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()

	must(err)
	must(json.NewEncoder(w).Encode(lists))
}