package main

import (
	"fmt"
	"go-book/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(router *mux.Router) {
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", middlewares.LoggingMiddleware(router))
}

