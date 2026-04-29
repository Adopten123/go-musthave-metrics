package main

import (
	"net/http"

	"github.com/Adopten123/go-musthave-metrics/internal/handler"
	"github.com/Adopten123/go-musthave-metrics/internal/storage"
)

func main() {
	s := storage.NewMemStorage()

	http.HandleFunc("/update/", handler.UpdateHandler(s))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
