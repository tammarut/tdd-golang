package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 100
}
func (s *InMemoryPlayerStore) RecordWin(name string) {
}

func main() {
	saber := InMemoryPlayerStore{}
	server := &PlayerServer{&saber}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000, %v", err)
	}
}
