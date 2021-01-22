package server

import (
	"logbyte/src/types"
	"net/http"
)

func auth(cfg *types.ConfigServer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				http.Error(w, http.StatusText(403), 403)
				return
			}
			if auth != cfg.Auth {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
