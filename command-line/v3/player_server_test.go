package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)


func TestPlayerServer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Maxu": 20,
            "Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)
	t.Run("return player's score", func(t *testing.T) {
		request := newGetScoreRequest("Maxu")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("get 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Nancy")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
		winCalls: nil,
	}
	server := NewPlayerServer(&store)
	t.Run("it returns accepted on post request", func(t *testing.T) {
		player := "nancy"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
            t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
        }
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}


// 测试返回成员列表 json
func TestLeague(t *testing.T) {
	
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := League{
            {"Cleo", 32},
            {"Chris", 20},
            {"Tiest", 14},
        }
		store := StubPlayerStore{nil, nil, wantedLeague}
    	server := NewPlayerServer(&store)

		request:= newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		got := getLeagueFromResponse(t, response.Body)

		assertLeague(t, got, wantedLeague)
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
func newPostScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}


func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}



func assertStatus(t *testing.T, got, want int) {
	t.Helper()
    if got != want {
        t.Errorf("did not get correct status, got %d, want %d", got, want)
    }
}

func getLeagueFromResponse(t *testing.T,body io.Reader) (league League) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return
}

func assertLeague(t *testing.T, got, want League) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func newLeagueRequest() *http.Request {
    req, _ := http.NewRequest(http.MethodGet, "/league", nil)
    return req
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
    if response.Header().Get("content-type") != want {
        t.Errorf("response did not have content-type of %s, got %v", want, response.HeaderMap)
    }
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}