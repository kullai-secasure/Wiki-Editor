package middleware

import (
	"context"
	"crypto/subtle"
	"net/http"
	"os"
	"strings"
)

type contextKey string

const UserKey contextKey = "user"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Reading wiki pages is intentionally public. Every other route
		// (edits, API) requires a validated editor token.
		if r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/wiki/") {
			next.ServeHTTP(w, r)
			return
		}

		expected := os.Getenv("EDITOR_TOKEN")
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if expected == "" || subtle.ConstantTimeCompare([]byte(token), []byte(expected)) != 1 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, "editor")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
