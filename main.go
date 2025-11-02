package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := openDB()
	if err != nil { log.Fatal(err) }
	defer db.Close()

	api := restRoutes(db)

	mux := http.NewServeMux()
	mux.Handle("/", api) // serves /api/todos, etc.

	log.Println("🚀 REST server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
