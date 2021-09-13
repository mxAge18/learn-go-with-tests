package main

import (
	"log"
	"net/http"
	"os"
)

/*
v2 版本中的问题：
写入数据返回到文件的开头，然后写入新的数据，但是如果新的数据比之前的数据要小 会出现问题
尽管目前没有删除功能，肯定是写入更多的数据。
v3 中新建tape.go
*/

const dbFileName = "game.db.json"
func main() {
	osfile, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }
	store := NewFileSystemStore(osfile)
	
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := NewPlayerServer(store)
	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}