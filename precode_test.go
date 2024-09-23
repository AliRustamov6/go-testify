package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) // здесь создан запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверка статуса ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// проверка тело ответа
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body) // убедимся что тело не пустое

	cafe := strings.Split(body, ",")
	assert.Len(t, cafe, 2) // ожидаем 2 кафе
}

func TestMainHandlerInvalidCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=3&city=krasnodar", nil) // здесь создан запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверка статуса ответа
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// проверка тело ответа
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body) // оджидается сообшение об ошибке
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь создан запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверка статуса ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// проверка тело ответа
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body) // убедимся что тело не пустое

	cafe := strings.Split(body, ",")
	assert.Len(t, cafe, totalCount) // ожидаем 4 кафе
}
