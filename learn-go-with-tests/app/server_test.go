package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestPlayerServer(t *testing.T) {
// 	t.Run("PlayerServer", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
// 		response := httptest.NewRecorder()
// 		PlayerServer(response, request)
// 		got := response.Body.String()
// 		want := "10"
// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

// func TestGetPlayers(t *testing.T) {
// 	t.Run("GetPlayerScore", func(t *testing.T) {
// 		request := newGetScoreRequest("Pepper")
// 		response := httptest.NewRecorder()
// 		PlayerServer(response, request)
// 		assertResponseBody(t, response.Body.String(), "20")
// 	})
// }

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response boy got %q, want %q", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response status code %q, want %q", got, want)
	}
}

func TestGetPlayer(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: []string{},
	}
	// server := &PlayerServer{store: store}
	server := NewPlayerServer(store)

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{"Pepper"}, "20"},
		{"B", args{"Floyd"}, "10"},
		{"C", args{"Apollo"}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetScoreRequest(tt.args.name)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			assertResponseBody(t, response.Body.String(), tt.want)
		})
	}
}

func TestLeague(t *testing.T) {
	store := &StubPlayerStore{}
	// server := &PlayerServer{store: store}
	server := NewPlayerServer(store)

	t.Run("league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
	})
}
