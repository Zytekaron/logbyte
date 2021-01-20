package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"notifs/src/notifs"
	"notifs/src/types"
	"strconv"
)

func Start(cfg *types.ConfigServer) {
	handler := router(cfg)
	err := http.ListenAndServe(":"+strconv.Itoa(cfg.Port), handler)
	log.Fatal(err)
}

func router(cfg *types.ConfigServer) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(auth(cfg))

	r.Get("/{id}", notifs.Get)
	r.Post("/", notifs.Post)
	//r.Put("/", notifs.Put)
	r.Patch("/", notifs.Patch)
	r.Delete("/", notifs.Delete)

	return r
}
