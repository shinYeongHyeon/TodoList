package api

import (
	"encoding/json"
	"github.com/shinYeongHyeon/TodoList/db"
	"net/http"
	"github.com/gorilla/mux"
)

// TodoListAPI Hello world
func TodoListAPI() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)

	return router
}

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, _ := db.GetTodoLists()
	json.NewEncoder(w).Encode(lists)
}
