package main

import (
	"fmt"
	"net/http"
)

type StorageGame interface {
	GetGameYearRelease(game string) int
	RegisterReleasedYear(game string)
}

type StorageGameInMemory struct{}

type ServerGame struct {
	storage StorageGame
}

func (s *ServerGame) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.registerReleasedYear(w)
	case http.MethodGet:
		s.showReleasedYear(w, r)
	}

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

func (s *StorageGameInMemory) RegisterReleasedYear(game string) {}

func (s *ServerGame) showReleasedYear(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/games/"):]
	releasedYear := s.storage.GetGameYearRelease(game)

	if releasedYear == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, releasedYear)
}

func (s *ServerGame) registerReleasedYear(w http.ResponseWriter) {
	s.storage.RegisterReleasedYear("MortalKombat")
	w.WriteHeader(http.StatusAccepted)
}
