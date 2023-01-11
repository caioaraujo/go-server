package main

import (
	"fmt"
	"net/http"
)

func ServerGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1990")
}
