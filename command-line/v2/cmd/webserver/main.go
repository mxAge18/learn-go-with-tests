package main

import (
	"log"
	"net/http"
	"os"
	"github.com/mxage18/learn-go-with-tests/command-line/v2"
)


const dbFileName = "game.db.json"
func main() {
	osfile, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }
	store, err := poker.NewFileSystemStore(osfile)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := poker.NewPlayerServer(store)
	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}