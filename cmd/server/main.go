package main

import (
	"log"
	"net/http"

	"wikiflex/internal/handlers"
	"wikiflex/internal/middleware"
	"wikiflex/internal/storage"
)

func main() {
	db := storage.NewPostgresDB()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/wiki/", handlers.ViewPage(db))
	mux.HandleFunc("/edit/", handlers.EditPage(db))
	mux.HandleFunc("/api/pages", handlers.ListPages(db))

	wrapped := middleware.Logger(middleware.SecurityHeaders(middleware.Auth(mux)))
	log.Println("WikiFlex started on :8080")
	log.Fatal(http.ListenAndServe(":8080", wrapped))
}
