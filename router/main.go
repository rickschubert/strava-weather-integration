package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/rick", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Guten Tag")
	})

	router.Get("/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My second route")
	})
	return router
}
