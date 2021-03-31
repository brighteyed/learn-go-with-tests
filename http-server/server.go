package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ErrPlayerNotFound = errors.New("player not found")

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score, err := p.store.GetPlayerScore(player)
	if err == ErrPlayerNotFound {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
