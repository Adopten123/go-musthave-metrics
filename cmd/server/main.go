package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/Adopten123/go-musthave-metrics/internal/handler"
	"github.com/Adopten123/go-musthave-metrics/internal/logger"
	"github.com/Adopten123/go-musthave-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	var flagRunAddr string
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		flagRunAddr = envRunAddr
	}

	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}

	s := storage.NewMemStorage()
	r := chi.NewRouter()

	r.Use(logger.RequestLogger)

	r.Post("/update/{type}/{name}/{value}", handler.UpdateHandler(s))
	r.Get("/value/{type}/{name}", handler.ValueHandler(s))
	r.Get("/", handler.AllMetricsHandler(s))

	logger.Log.Info("Starting server", zap.String("address", flagRunAddr))

	err := http.ListenAndServe(flagRunAddr, r)
	if err != nil {
		logger.Log.Fatal("Server crashed", zap.Error(err))
	}
}
