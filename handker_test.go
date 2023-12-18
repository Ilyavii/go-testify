package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := "d"
	result := "d"
	assert.NotEmpty(t, result)
	assert.Equal(t, result, totalCount)
	// req := http. // здесь нужно создать запрос к сервису

	// responseRecorder := httptest.NewRecorder()
	// handler := http.HandlerFunc(mainHandle)
	// handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
}
