package handler

import (
	"fmt"
	"net/http"
)

func AllMetricsHandler(s MetricStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		html := "<html><body><h1>Метрики</h1><ul>"

		for name, val := range s.GetAllGauges() {
			html += fmt.Sprintf("<li>%s: %f</li>", name, val)
		}
		for name, val := range s.GetAllCounters() {
			html += fmt.Sprintf("<li>%s: %d</li>", name, val)
		}

		html += "</ul></body></html>"

		w.Write([]byte(html))
	}
}
