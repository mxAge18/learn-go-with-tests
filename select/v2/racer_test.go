package v2

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns an error if a server doesn't respond within 20 * time.Millisecond", func(t *testing.T) {
		serverA := makeDelayedServer(26 * time.Millisecond)
		serverB := makeDelayedServer(25 * time.Millisecond)
	
		defer serverA.Close()
		defer serverB.Close()
	
		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 20 * time.Millisecond)
	
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
	t.Run("two normal request return a fast url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got,_ := Racer(slowURL, fastURL)
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}