package server

import (
	"net/http"
	"os"
)

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, http.StatusText(403), 403)
			return
		}
		if auth != os.Getenv("ZYTEKARON_AUTH") {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		next.ServeHTTP(w, r)
	})
}
