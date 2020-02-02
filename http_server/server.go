package main

import (
	"fmt"
	"net/http"
	"strings"
)

type StubPlayerStore struct {
	scores map[string]int
}
type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	scores := p.store.GetPlayerScore(player)
	if scores == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	scores := s.scores[name]
	return scores
}
