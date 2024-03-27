package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест, если запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое

func TestMainHandlerWhenRequestIsCorrectAndResponseHasValue(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	http.HandlerFunc(mainHandle).ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

// Тест, если город, который передаётся в параметре `city`, не поддерживается. Сервис возвращает код ответа 400 и ошибку `wrong city value` в теле ответа
func TestMainHandlerWhenCityNotExist(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=saint-petersburg", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	bodyStr := responseRecorder.Body.String()
	status := responseRecorder.Code

	expectedValue := "wrong city value"

	require.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, expectedValue, bodyStr)
}

// Тест, если в параметре `count` указано больше, чем есть всего, должны вернуться все доступные кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=999&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, totalCount, len(list))
}
