package main

import (
	"log"
	"net/http"
	"os"
)

/*
在 RecordWin 中，我们有行 json.NewEncoder(f.database).Encode(f.league)。
我们不需要在每次编写代码时创建一个新的编码器，我们可以在构造函数中初始化一个编码器并使用它。
在我们的类型中存储对编码器的引用。
*/

const dbFileName = "game.db.json"
func main() {
	osfile, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }
	store, err := NewFileSystemStore(osfile)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := NewPlayerServer(store)
	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}