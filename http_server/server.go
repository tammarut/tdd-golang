package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	router.HandleFunc("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := r.URL.Path[len("/players/"):]

		switch r.Method {
		case http.MethodPost:
			p.processWin(w, r, player)
		case http.MethodGet:
			p.showScore(w, r, player)
		}
	}))

	router.ServeHTTP(w, r)

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
