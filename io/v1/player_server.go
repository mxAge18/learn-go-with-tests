package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
    GetLeague() []Player
}
type PlayerServer struct {
    store PlayerStore
    http.Handler
}

type Player struct {
    Name string
    Wins int
}
const jsonContentType = "application/json"
func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)
    p.store = store

    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.LeagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.PlayersHandler))

    p.Handler = router
    return p
}


// 处理League
func (p *PlayerServer) LeagueHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", jsonContentType)
    json.NewEncoder(w).Encode(p.store.GetLeague())
    w.WriteHeader(http.StatusOK)
}



// 处理players相关请求
func (p *PlayerServer) PlayersHandler(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]
    switch r.Method {
    case http.MethodGet:
        p.showScore(w, player)
    case http.MethodPost:
        p.processWin(w, player)
    }
}

func (p *PlayerServer) showScore(w http.ResponseWriter,player string) {
    score := p.store.GetPlayerScore(player)
    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }
    
    fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
    p.store.RecordWin(player)
    w.WriteHeader(http.StatusAccepted)
   
}

