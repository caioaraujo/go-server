package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGame(t *testing.T) {
	t.Run("return Super Mario World", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/games/SuperMarioWorld", nil)
		response := httptest.NewRecorder()

		ServerGame(response, request)

		obtained := response.Body.String()
		expected := "1990"

		if obtained != expected {
			t.Errorf("obtained '%s', expected '%s'", obtained, expected)
		}
	})
}
