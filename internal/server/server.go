package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
)

func NewServer() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	apiRouter := chi.NewRouter()
	v1Router := chi.NewRouter()

	v1Router.Group(func(r chi.Router) {
		r.Post("/auth/login", Login)
	})

	r.Mount("/api", apiRouter)
	apiRouter.Mount("/v1", v1Router)
	return r
}
