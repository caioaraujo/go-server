package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStorageGame struct {
	releasedYears        map[string]int
	registerReleasedYear []string
}

func (e *MockStorageGame) GetGameYearRelease(game string) int {
	releasedYear := e.releasedYears[game]
	return releasedYear
}

func (e *MockStorageGame) RegisterReleasedYear(game string) {
	e.registerReleasedYear = append(e.registerReleasedYear, game)
}

func TestGetGame(t *testing.T) {
	storage := MockStorageGame{
		map[string]int{
			"SuperMarioWorld": 1990,
			"SuperMetroid":    1994,
		},
		nil,
	}
	server := &ServerGame{&storage}

	t.Run("return Super Mario World", func(t *testing.T) {
		request := newRequestGetYearRelease("SuperMarioWorld")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseStatusCode(t, response.Code, http.StatusOK)
		checkRequestBody(t, response.Body.String(), "1990")
	})

	t.Run("return Super Metroid", func(t *testing.T) {
		request := newRequestGetYearRelease("SuperMetroid")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseStatusCode(t, response.Code, http.StatusOK)
		checkRequestBody(t, response.Body.String(), "1994")
	})

	t.Run("return 404 for not found game", func(t *testing.T) {
		request := newRequestGetYearRelease("GoldenAxe")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStorageReleasedYear(t *testing.T) {
	storage := MockStorageGame{
		map[string]int{},
		nil,
	}
	server := &ServerGame{&storage}

	t.Run("return 'accepted' status for method POST calls", func(t *testing.T) {
		request := newRequestRegisterReleasedYearPost("SuperMarioWorld")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseStatusCode(t, response.Code, http.StatusAccepted)

		if len(storage.registerReleasedYear) != 1 {
			t.Errorf("after checking %d calls to RegisterReleasedYear, expected %d", len(storage.registerReleasedYear), 1)
		}
	})

}

func newRequestGetYearRelease(game string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/games/%s", game), nil)
	return request
}

func newRequestRegisterReleasedYearPost(game string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/games/%s", game), nil)
	return request
}

func checkRequestBody(t *testing.T, obtained, expected string) {
	t.Helper()
	if obtained != expected {
		t.Errorf("invalid request body, obtained '%s', expected '%s'", obtained, expected)
	}
}

func checkResponseStatusCode(t *testing.T, obtained, expected int) {
	t.Helper()
	if obtained != expected {
		t.Errorf("could not find the expected HTTP status code. Obtained %d, expected %d", obtained, expected)
	}
}
