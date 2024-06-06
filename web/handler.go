package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"prostoTak"
)

type Handler struct {
	*chi.Mux

	blog prostoTak.Store
}

func NewHandler(blog prostoTak.Store) *Handler {
	h := &Handler{Mux: chi.NewMux(), blog: blog}

	h.Use(middleware.Logger)
	h.Route("/auth", func(r chi.Router) {
		//r.Post("/sign-up", h.signUp())
		//r.Post("/sign-in", h.signIn())
	})

	h.Route("/blog", func(r chi.Router) {
		r.Get("/", h.HelloPage())
		r.Get("/list", h.BlogList())
		r.Get("/new", h.BlogCreate())
		r.Post("/", h.BlogSave())
		r.Get("/{id}", h.Blog())
		r.Group(func(r chi.Router) {
			r.Use(middleware.BasicAuth("MyRealm", map[string]string{
				"bob111":   "password1",
				"alice111": "password2",
			}))

			r.Post("/{id}/delete", h.BlogDelete())
			r.Post("/{id}/up", h.BlogUp())
			r.Post("/{id}/down", h.BlogDown())
		})
	})

	return h
}
