package main

import (
	"log"
	"net/http"
)
// 要求有一个叫 /league（联盟）的新的端点（endpoint），
// 它可以返回一个玩家清单。她想让它返回一个 JSON 格式的数据。
type InMemoryPlayerStore struct{
	store map[string]int
	league []Player
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return i.store[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
func (i *InMemoryPlayerStore) GetLeague() []Player {
	for name, wins := range i.store {
		i.league = append(i.league, Player{name, wins})
	}
    return i.league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
		league: make([]Player, 0),
	}
}
func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}