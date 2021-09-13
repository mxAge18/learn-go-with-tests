package main

import (
	"log"
	"net/http"
	"os"
)
const dbFileName = "game.db.json"
func main() {
	osfile, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }
	server := NewPlayerServer(NewFileSystemStore(osfile))
	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}