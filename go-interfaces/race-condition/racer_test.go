package condition

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		serverA := makeDelayedServer(20 * time.Millisecond)
		serverB := makeDelayedServer(0 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()
		want := serverB.URL
		got, err := Racer(serverA.URL, serverB.URL)

		if err != nil {
			t.Errorf("didn't expect an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
