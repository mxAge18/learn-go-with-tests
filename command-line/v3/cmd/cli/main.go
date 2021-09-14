package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	poker "github.com/mxage18/learn-go-with-tests/command-line/v3"
)

const dbFileName = "game.db.json"
func main() {
	
	store, err := poker.FileSystemStoreFromFile(dbFileName)
	
	fmt.Println(GetAppPath())
  	if err != nil {
       log.Fatal(err)
    }
	
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}