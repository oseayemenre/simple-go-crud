package routes

import (
	"github.com/go-chi/chi"
	"github.com/oseayemenre/go_crud_scratch/internal/handlers"
	"github.com/oseayemenre/go_crud_scratch/internal/middlewares"
	"github.com/oseayemenre/go_crud_scratch/internal/types"
)

type DB struct {
	*types.ApiConfig
}

func (d *DB) HandleRoutes(r *chi.Mux) {

	r.Use(middlewares.CorsMiddleware)

	v1 := chi.NewRouter()

	r.Mount("/v1", v1)

	v1.Get("/healthcheck", handlers.HealthCheck)
}
