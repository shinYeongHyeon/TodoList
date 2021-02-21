package main

import (
	"github.com/shinYeongHyeon/TodoList/server"
	"log"
)

func main() {
	if err := server.ListenAndServe(server.Config {
		Address: ":8080",
		DatabaseURL: "host=localhost port=5432 user=postgres dbname=go-todo sslmode=disable",
	}); err != nil {
		log.Fatalln(err)
	}
}