package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type StorageGame interface {
	GetGameYearRelease(game string) int
	RegisterReleasedYear(game string, releasedYear int)
}

func NewStorageGameInMemory() *StorageGameInMemory {
	return &StorageGameInMemory{map[string]int{}}
}

type StorageGameInMemory struct {
	storage map[string]int
}

type ServerGame struct {
	storage StorageGame
}

func (s *ServerGame) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/games/"):]

	switch r.Method {
	case http.MethodPost:
		var res map[string]string
		json.NewDecoder(r.Body).Decode(&res)
		releasedYear := res["releasedYear"]
		releasedYearAsInt, _ := strconv.Atoi(releasedYear)
		s.registerReleasedYear(w, game, releasedYearAsInt)
	case http.MethodGet:
		s.showReleasedYear(w, game)
	}

}

func (s *StorageGameInMemory) GetGameYearRelease(game string) int {
	return s.storage[game]
}

func (s *StorageGameInMemory) RegisterReleasedYear(game string, releasedYear int) {
	s.storage[game] = releasedYear
}

func (s *ServerGame) showReleasedYear(w http.ResponseWriter, game string) {
	releasedYear := s.storage.GetGameYearRelease(game)

	if releasedYear == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, releasedYear)
}

func (s *ServerGame) registerReleasedYear(w http.ResponseWriter, game string, releasedYear int) {
	s.storage.RegisterReleasedYear(game, releasedYear)
	w.WriteHeader(http.StatusAccepted)
}
