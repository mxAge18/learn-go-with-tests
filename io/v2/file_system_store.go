package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
	league League
}

func NewFileSystemStore(db io.ReadWriteSeeker) *FileSystemStore {
	db.Seek(0, 0)
    league, _ := NewLeague(db)
	return &FileSystemStore{
        database:db,
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
    // league := f.GetLeague()
	// for index, player := range league {
	// 	if player.Name == name {
	// 		league[index].Wins++
	// 		// player.Wins++ 当你在一个切片上取值时，
	// 		// 将返回当前循环的索引（我们示例中的 i）和该索引中的元素的副本。
	// 		// 更改副本 Wins 的值不会对我们迭代的 league 产生任何影响。
	// 		// 因此，我们需要通过使用 league[i] 来获取对实际值的引用，然后更改该值。
	// 	}
	// }
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}
	f.database.Seek(0,0)
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