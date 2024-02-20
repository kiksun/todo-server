package main

import (
	"log"
	"net/http"

	"github.com/kiksun/todo-server/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HelloHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
