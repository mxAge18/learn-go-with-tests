package main

import (
	"log"
	"net/http"
	"github.com/mxage18/learn-go-with-tests/command-line/v3"
)


const dbFileName = "game.db.json"
func main() {
	store, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	server := poker.NewPlayerServer(store)

	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}