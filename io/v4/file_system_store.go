package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemStore实现了接口PlayerStore
type FileSystemStore struct {
	database *json.Encoder
	league League
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

    info, err := file.Stat()

    if err != nil {
        return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
    }

    if info.Size()==0 {
        file.Write([]byte("[]"))
        file.Seek(0, 0)
    }

    return nil
}

func NewFileSystemStore(db *os.File) (*FileSystemStore, error) {
	err := initialisePlayerDBFile(db)

	if err != nil {
		return nil,  fmt.Errorf("problem initialising player db file, %v", err)
	}

    league, err := NewLeague(db)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", db.Name(), err)
	}

	return &FileSystemStore{
        database:json.NewEncoder(&tape{db}),
        league:league,
    }, nil
}

func (f *FileSystemStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
    player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}
	f.database.Encode(f.league)
}



// 抽象出League,并写一个find方法
type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}