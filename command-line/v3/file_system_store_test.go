package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("didnt expect an error but got one, %v", err)
    }
}
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
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase	:= createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)

		assertNoError(t, err)
		got := store.GetLeague()

		want := League{
			{"Chris", 33},
            {"Cleo", 10},
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
		
		store, err := NewFileSystemStore(database)
		// store := FileSystemStore{database}
		assertNoError(t, err)
		assertScoreEquals(t, store.GetPlayerScore("Chris"), 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase	:= createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		player := "Chris"
		store, err := NewFileSystemStore(database)
		// store := FileSystemStore{database}
		assertNoError(t, err)
		store.RecordWin(player)
		assertScoreEquals(t, store.GetPlayerScore(player), 34)
	})
	
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase	:= createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		player := "Maxu"
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		store.RecordWin(player)
		assertScoreEquals(t, store.GetPlayerScore(player), 1)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()
	
		_, err := NewFileSystemStore(database)
	
		assertNoError(t, err)
	})
}
