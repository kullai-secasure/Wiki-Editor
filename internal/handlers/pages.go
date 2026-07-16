package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"wikiflex/internal/renderer"
	"wikiflex/internal/storage"
)

func ViewPage(db *storage.PostgresDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/wiki/")
		page, err := db.GetPageBySlug(slug)
		if err != nil {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
		html, err := renderer.RenderPage(page)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	}
}

func EditPage(db *storage.PostgresDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Edit page handler"))
	}
}

func ListPages(db *storage.PostgresDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}
}
