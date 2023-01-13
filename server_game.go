package main

import (
	"fmt"
	"net/http"
)

type StorageGame interface {
	GetGameYearRelease(game string) int
}

type ServerGame struct {
	storage StorageGame
}

func (s *ServerGame) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/games/"):]
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, s.storage.GetGameYearRelease(game))
}

func GetGameYearRelease(game string) string {
	if game == "SuperMarioWorld" {
		return "1990"
	}

	if game == "SuperMetroid" {
		return "1994"
	}

	return ""
}
