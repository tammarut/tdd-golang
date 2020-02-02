package main

import (
	"log"
	"net/http"
)

type SaberScores struct{}

func (s *SaberScores) GetPlayerScore(name string) int {
	return 100
}
func (s *SaberScores) RecordWin(name string) {
}

func main() {
	saber := SaberScores{}
	server := &PlayerServer{&saber}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000, %v", err)
	}
}
