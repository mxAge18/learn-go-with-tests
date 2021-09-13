package poker

import (
	"testing"
)

type StubPlayerStore struct {
    scores map[string]int
	winCalls []string
	league League // 存储player的列表
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}
func (s *StubPlayerStore) RecordWin(name string) {
    s.winCalls = append(s.winCalls, name)
}

// 获取league相关数据
func (s *StubPlayerStore) GetLeague() League {
    return League{
        {"Cleo", 32},
        {"Chris", 20},
        {"Tiest", 14},
    }
}


func AssertPlayerWin(t *testing.T, playerStore *StubPlayerStore, want string) {
	if len(playerStore.winCalls) < 1 {
        t.Fatal("expected a win call but didn't get any")
    }
	
	got := playerStore.winCalls[0]

	if got != want {
        t.Errorf("didn't record correct winner, got '%s', want '%s'", got, want)
    }
}
