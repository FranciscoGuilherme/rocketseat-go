package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Get("/horario", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintln(w, now)
	})

	r.Route("/api", func (r chi.Router) {
		r.Route("/v1", func (r chi.Router) {
			r.Get("/users", func (w http.ResponseWriter, r *http.Request) {})
			r.Get("/users/{id}", func (w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				fmt.Fprintln(w, id)
			})
			r.Get("/users/{id:[0-9]+}", func (w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				fmt.Fprintln(w, id)
			})
		})
	})

	r.Group(func (r chi.Router) {
		r.Use(middleware.BasicAuth("", map[string]string{
			"admin": "admin",
		}))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
