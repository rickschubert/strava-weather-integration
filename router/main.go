package router

import (
	"strava-weather-integration/router/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func setupMiddleware(router *chi.Mux) {
	router.Use(middleware.Logger)
}

func setupRoutes(router *chi.Mux) {
	router.Get("/runs", handlers.GetRuns)
	router.Get("/auth/redirect", handlers.OAuth2Redirect)
	router.Get("/auth/authenticate", handlers.OAuth2Authenticate)
	router.NotFound(handlers.RouteNotFound)
}

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	setupMiddleware(router)
	setupRoutes(router)
	return router
}
