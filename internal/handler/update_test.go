package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Adopten123/go-musthave-metrics/internal/storage"
)

func TestUpdateHandler(t *testing.T) {
	s := storage.NewMemStorage()
	handler := UpdateHandler(s)

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
			name:         "Неверный метод (GET вместо POST)",
			method:       http.MethodGet,
			url:          "/update/gauge/Alloc/12.5",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "Неизвестный тип метрики",
			method:       http.MethodPost,
			url:          "/update/unknown/test/123",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Некорректное значение счетчика (строка)",
			method:       http.MethodPost,
			url:          "/update/counter/test/none",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Отсутствует значение метрики в URL",
			method:       http.MethodPost,
			url:          "/update/counter/test/",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, tc.url, nil)
			recorder := httptest.NewRecorder()

			handler(recorder, request)

			res := recorder.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expectedCode {
				t.Errorf("Ожидался статус %d, получен %d", tc.expectedCode, res.StatusCode)
			}
		})
	}
}
