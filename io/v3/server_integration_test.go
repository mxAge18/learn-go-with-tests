package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// store := InMemoryPlayerStore{store: map[string]int{}}
	db, closeDb := createTempFile(t, "")
	defer closeDb()
	store := FileSystemStore{database: db}
	server := NewPlayerServer(&store)
	player := "cheng"
	// 让cheng的得分为3
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := League{
            {"cheng", 3},
        }
		assertLeague(t, got, want)
	})

}