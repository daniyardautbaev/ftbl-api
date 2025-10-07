package router

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/daniyardautbaev/ftbl-api/internal/handlers"
)

func NewRouter() http.Handler {
    r := chi.NewRouter()

    r.Route("/teams", func(r chi.Router) {
        r.Post("/", handlers.CreateTeam)
        r.Get("/", handlers.ListTeams)
    })

    r.Route("/players", func(r chi.Router) {
        r.Post("/", handlers.CreatePlayer)
        r.Get("/", handlers.ListPlayers)
        r.Get("/{id}", handlers.GetPlayer)
        r.Put("/{id}", handlers.UpdatePlayer)
        r.Delete("/{id}", handlers.DeletePlayer)
    })

    // health
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    return r
}
