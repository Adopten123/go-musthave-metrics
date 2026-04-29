package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Adopten123/go-musthave-metrics/internal/storage"
)

func main() {
	s := storage.NewMemStorage()

	http.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 5 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		mType, mName, mValue := parts[2], parts[3], parts[4]

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
	})

	http.ListenAndServe(":8080", nil)
}
