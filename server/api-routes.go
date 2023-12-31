package server

import (
	"net/http"

	"github.com/DeanRTaylor1/deans-site/handlers"
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterApiRoutes() {
	apiRouter := chi.NewRouter()
	apiRouter.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			s.HealthCheck(w, r)
		})

		r.Post("/subscribe", func(w http.ResponseWriter, r *http.Request) {
			handlers.Subscribe(w, r, *s.Logger, s.MongoClient, &s.Validator, s.Config)
		})

		r.Get("/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
			email := r.URL.Query().Get("email")
			handlers.Unsubscribe(w, r, *s.Logger, s.MongoClient, &s.Validator, email)
		})

		r.Post("/contact", func(w http.ResponseWriter, r *http.Request) {
			handlers.PostContact(w, r, *s.Logger, s.MongoClient, &s.Validator, s.Config)
		})
	})

	s.Router.Mount("/", apiRouter)
}
