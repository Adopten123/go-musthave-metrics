package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateHandler(s MetricStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mType := chi.URLParam(r, "type")
		mName := chi.URLParam(r, "name")
		mValue := chi.URLParam(r, "value")

		if mName == "" || mValue == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch mType {
		case "counter":
			v, err := strconv.ParseInt(mValue, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			s.UpdateCounter(mName, v)
		case "gauge":
			v, err := strconv.ParseFloat(mValue, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			s.UpdateGauge(mName, v)
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	}
}
