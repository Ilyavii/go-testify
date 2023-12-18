package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrect(t *testing.T) {
	responseRecorder := TestRequest("/cafe?count=2&city=moscow")

	expectedCode := http.StatusOK
	expectedCount := 2

	receivedCount := strings.Split(responseRecorder.Body.String(), ",")

	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, expectedCode, responseRecorder.Code)
	assert.Len(t, receivedCount, expectedCount)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	responseRecorder := TestRequest("/cafe?count=10&city=moscow")

	expectedCode := http.StatusOK
	expectedCount := 4

	receivedCount := strings.Split(responseRecorder.Body.String(), ",")

	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, expectedCode, responseRecorder.Code)
	assert.Len(t, receivedCount, expectedCount)
}

func TestMainHandlerWhenIncorrectCity(t *testing.T) {
	responseRecorder := TestRequest("/cafe?count=2&city=kazan")

	expectedCode := http.StatusBadRequest

	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, expectedCode, responseRecorder.Code)
}
