package main

import "go-book/pkg/db"

func main() {
	db.InitDB()
	defer db.CloseDB()

	mux := RegisterRoutes()

	StartServer(mux)
}
