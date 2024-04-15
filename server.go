package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusCreated)
		return
	}

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	playerScore := p.store.GetPlayerScore(player)

	if playerScore == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, playerScore)
}
