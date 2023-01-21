package main

import (
	"bytes"
	"encoding/json"
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

func (e *MockStorageGame) RegisterReleasedYear(game string, releasedYear int) {
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

	t.Run("post released year for a specific game", func(t *testing.T) {
		game := "FinalFantasyVI"
		releasedYear := "1994"

		request := newRequestRegisterReleasedYearPost(game, releasedYear)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseStatusCode(t, response.Code, http.StatusAccepted)

		if len(storage.registerReleasedYear) != 1 {
			t.Errorf("after checking %d calls to RegisterReleasedYear, expected %d", len(storage.registerReleasedYear), 1)
		}

		if storage.registerReleasedYear[0] != game {
			t.Errorf("could not record the game correctly, obtained '%s', expected '%s'", storage.registerReleasedYear[0], game)
		}
	})

}

func TestPostAndGetReleasedYear(t *testing.T) {
	storage := NewStorageGameInMemory()
	server := ServerGame{storage}
	game := "SuperMarioWorld"
	releasedYear := "1990"

	server.ServeHTTP(httptest.NewRecorder(), newRequestRegisterReleasedYearPost(game, releasedYear))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newRequestGetYearRelease(game))
	checkResponseStatusCode(t, response.Code, http.StatusOK)

	checkRequestBody(t, response.Body.String(), releasedYear)
}

func newRequestGetYearRelease(game string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/games/%s", game), nil)
	return request
}

func newRequestRegisterReleasedYearPost(game string, releasedYear string) *http.Request {
	body, _ := json.Marshal(map[string]string{
		"releasedYear": releasedYear,
	})
	payload := bytes.NewBuffer(body)

	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/games/%s", game), payload)
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
