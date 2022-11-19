package mmain

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)
func routes() http.Handler{
	mux := chi.NewRouter()

	// Specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeders: []string{"Accept", "Authorization","Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowedCredentials: true,
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	return mux
}