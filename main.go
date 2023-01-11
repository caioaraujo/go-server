package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(ServerGame)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen in port 5000 %v", err)
	}
}
