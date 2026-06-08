package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// responseData хранит сведения об ответе.
type responseData struct {
	status int
	size   int
}

// loggingResponseWriter добавляет реализацию http.ResponseWriter.
type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

// Write перехватывает размер записанного ответа.
func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

// WriteHeader перехватывает код статуса ответа.
func (r *loggingResponseWriter) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.responseData.status = status
}

// RequestLogger — middleware-логер для входящих HTTP-запросов.
func RequestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		Log.Info("got incoming HTTP request",
			zap.String("uri", r.RequestURI),
			zap.String("method", r.Method),
			zap.Int("status", responseData.status),
			zap.Duration("duration", duration),
			zap.Int("size", responseData.size),
		)
	})
}
