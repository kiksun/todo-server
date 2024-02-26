package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiksun/todo-server/handlers"
)

type Todo struct {
	CommentID int
	ArticleID int
	content   string
}

func main() {
	todo1 := Todo{
		CommentID: 1,
		ArticleID: 2,
		content:   "aaaaaa",
	}
	todoList := []Todo{todo1, todo1}

	jsonData, err := json.Marshal(todoList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", jsonData)

	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler)

	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/{id:[0-9]+}",
		handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
