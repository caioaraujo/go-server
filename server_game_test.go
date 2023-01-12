package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStorageGame struct {
	yearReleases map[string]int
}

func (e *MockStorageGame) GetGameYearRelease(game string) int {
	yearRelease := e.yearReleases[game]
	return yearRelease
}

func TestGetGame(t *testing.T) {
	storage := MockStorageGame{
		map[string]int{
			"SuperMarioWorld": 1990,
			"SuperMetroid":    1994,
		},
	}
	server := &ServerGame{&storage}

	t.Run("return Super Mario World", func(t *testing.T) {
		request := newRequestGetYearRelease("SuperMarioWorld")
		response := httptest.NewRecorder()

		server.ServerHttp(response, request)

		checkRequestBody(t, response.Body.String(), "1990")
	})

	t.Run("return Super Metroid", func(t *testing.T) {
		request := newRequestGetYearRelease("SuperMetroid")
		response := httptest.NewRecorder()

		server.ServerHttp(response, request)

		checkRequestBody(t, response.Body.String(), "1994")
	})
}

func newRequestGetYearRelease(game string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/games/%s", game), nil)
	return request
}

func checkRequestBody(t *testing.T, obtained, expected string) {
	t.Helper()
	if obtained != expected {
		t.Errorf("invalid request body, obtained '%s', expected '%s'", obtained, expected)
	}
}
