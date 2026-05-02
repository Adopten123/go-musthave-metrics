package main

import (
	"flag"
	"net/http"

	"github.com/Adopten123/go-musthave-metrics/internal/handler"
	"github.com/Adopten123/go-musthave-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	var flagRunAddr string
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")

	flag.Parse()

	s := storage.NewMemStorage()
	r := chi.NewRouter()

	r.Post("/update/{type}/{name}/{value}", handler.UpdateHandler(s))
	r.Get("/value/{type}/{name}", handler.ValueHandler(s))
	r.Get("/", handler.AllMetricsHandler(s))

	err := http.ListenAndServe(flagRunAddr, r)

	if err != nil {
		panic(err)
	}
}
