package main

import (
	"io/ioutil"
	"os"
	"testing"
)


func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removefile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removefile
}
// 测试数据存储到文件系统
func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase	:= createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store := NewFileSystemStore(database)
		// store := FileSystemStore{database}

		got := store.GetLeague()

		want := League{
			{"Cleo", 10},
            {"Chris", 33},
		}
		assertLeague(t, got, want)
		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		
		database, cleanDatabase	:= createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		
		store := NewFileSystemStore(database)
		assertScoreEquals(t, store.GetPlayerScore("Chris"), 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase	:= createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		player := "Chris"
		store:= NewFileSystemStore(database)
		store.RecordWin(player)
		assertScoreEquals(t, store.GetPlayerScore(player), 34)
	})
	
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase	:= createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		player := "Maxu"
		store := NewFileSystemStore(database)

		store.RecordWin(player)
		assertScoreEquals(t, store.GetPlayerScore(player), 1)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

