package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ValueHandler(s MetricStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mType := chi.URLParam(r, "type")
		mName := chi.URLParam(r, "name")

		switch mType {
		case "gauge":
			val, ok := s.GetGauge(mName)
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(strconv.FormatFloat(val, 'f', -1, 64)))
		case "counter":
			val, ok := s.GetCounter(mName)
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(strconv.FormatInt(val, 10)))
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
