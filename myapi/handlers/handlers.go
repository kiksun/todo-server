package handlers

import (
	"io"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		log.Println("Get")
	}

	if req.Method == http.MethodPost {
		log.Println("Post")
	}
	if req.Method == http.MethodPut {
		log.Println("Put")
	}
	io.WriteString(w, "Hello, world!\n")
}
