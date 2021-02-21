package api

import (
	"encoding/json"
	"github.com/shinYeongHyeon/TodoList/db"
	"log"
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
	lists, err := db.GetTodoLists()
	if err != nil {
		log.Println("internal err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error string
		} {
			Error: "Internal Error",
		})
	}
	json.NewEncoder(w).Encode(lists)
}
