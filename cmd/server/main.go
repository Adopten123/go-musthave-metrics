package main

import (
	"net/http"

	"github.com/Adopten123/go-musthave-metrics/internal/handler"
	"github.com/Adopten123/go-musthave-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	s := storage.NewMemStorage()

	r := chi.NewRouter()

	r.Post("/update/{type}/{name}/{value}", handler.UpdateHandler(s))
	r.Get("/value/{type}/{name}", handler.ValueHandler(s))
	r.Get("/", handler.AllMetricsHandler(s))

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
