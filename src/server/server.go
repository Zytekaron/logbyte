package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"logbyte/src/endpoints"
	"logbyte/src/types"
	"net/http"
	"strconv"
)

func Start(cfg *types.ConfigServer) {
	handler := router(cfg)

	log.Println("Listening on", cfg.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(cfg.Port), handler)
	log.Fatal(err)
}

func router(cfg *types.ConfigServer) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(auth(cfg))

	r.Get("/{id}", endpoints.Get)
	r.Get("/", endpoints.Batch)
	r.Post("/", endpoints.Post)
	r.Patch("/{id}", endpoints.Patch)
	r.Delete("/{id}", endpoints.Delete)

	return r
}
