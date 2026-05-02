package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Adopten123/go-musthave-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
)

func TestUpdateHandler(t *testing.T) {
	s := storage.NewMemStorage()
	h := UpdateHandler(s)

	tests := []struct {
		name         string
		method       string
		url          string
		expectedCode int
	}{
		{
			name:         "Успешное обновление gauge",
			method:       http.MethodPost,
			url:          "/update/gauge/Alloc/12.5",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Успешное обновление counter",
			method:       http.MethodPost,
			url:          "/update/counter/PollCount/1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Некорректное значение счетчика",
			method:       http.MethodPost,
			url:          "/update/counter/test/none",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Отсутствует значение",
			method:       http.MethodPost,
			url:          "/update/counter/test/",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, tc.url, nil)
			recorder := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Post("/update/{type}/{name}/{value}", h)

			r.ServeHTTP(recorder, request)

			res := recorder.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expectedCode {
				t.Errorf("Кейс '%s': ожидался статус %d, получен %d", tc.name, tc.expectedCode, res.StatusCode)
			}
		})
	}
}
