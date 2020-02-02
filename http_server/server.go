package main

import (
	"fmt"
	"net/http"
	"strings"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r, player)
	case http.MethodGet:
		p.showScore(w, r, player)
	}
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	scores := s.scores[name]
	return scores
}
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request, player string) {
	scores := p.store.GetPlayerScore(player)
	if scores == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, scores)
}
