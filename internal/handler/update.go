package handler

import (
	"net/http"
	"strconv"

	"github.com/Adopten123/go-musthave-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
)

func UpdateHandler(s *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mType := chi.URLParam(r, "type")
		mName := chi.URLParam(r, "name")
		mValueStr := chi.URLParam(r, "value")

		switch mType {
		case "gauge":
			val, err := strconv.ParseFloat(mValueStr, 64)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			s.UpdateGauge(mName, val)
		case "counter":
			val, err := strconv.ParseInt(mValueStr, 10, 64)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			s.UpdateCounter(mName, val)
		default:
			http.Error(w, "Bad Request: unknown type", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
