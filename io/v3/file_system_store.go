package main

import (
	"encoding/json"
	"io"
	"os"
)

// FileSystemStore实现了接口PlayerStore
type FileSystemStore struct {
	database io.Writer
	league League
}

func NewFileSystemStore(db *os.File) *FileSystemStore {
	db.Seek(0, 0)
    league, _ := NewLeague(db)
	return &FileSystemStore{
        database: &tape{db},
        league:league,
    }
}

func (f *FileSystemStore) GetLeague() League {
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
	json.NewEncoder(f.database).Encode(f.league)
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