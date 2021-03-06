package main

import (
	"log"
	"net/http"
)
type InMemoryPlayerStore struct{
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return i.store[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}
func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}