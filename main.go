package main

import (
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func main() {
	server := &ServerGame{NewStorageGameInMemory()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen in port 5000 %v", err)
	}
}
