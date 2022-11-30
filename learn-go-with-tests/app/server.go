package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// PlayerServer handles player requests
// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := r.URL.Path[len("/players/"):]
// 	fmt.Fprint(w, GetPlayerScore(player))
// }

// GetPlayerScore get player score by name
// func GetPlayerScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}
// 	if name == "Floyd" {
// 		return "10"
// 	}
// 	return ""
// }

// Player .
type Player struct {
	Name string
	Wins int
}

// PlayerStore of abstract
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWinPlayer(name string)
	GetLeague() League
}

// StubPlayerStore store player score with map
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

// GetPlayerScore get player score by name
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

// RecordWinPlayer records player which has won
func (s *StubPlayerStore) RecordWinPlayer(name string) {
	s.winCalls = append(s.winCalls, name)
}

// GetLeague return player's league
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// MemoryPlayerStore store player score in memory
type MemoryPlayerStore struct {
	store map[string]int
}

// NewMemoryPlayerStore create a MemoryPlayerStore instance
func NewMemoryPlayerStore() *MemoryPlayerStore {
	return &MemoryPlayerStore{map[string]int{}}
}

// GetPlayerScore get player score by name
func (m *MemoryPlayerStore) GetPlayerScore(name string) int {
	return m.store[name]
}

// RecordWinPlayer records player which has won
func (m *MemoryPlayerStore) RecordWinPlayer(name string) {
	m.store[name]++
}

// GetLeague return player's league
func (m *MemoryPlayerStore) GetLeague() []Player {
	league := make([]Player, 0)
	for k, v := range m.store {
		league = append(league, Player{Name: k, Wins: v})
	}
	return league
}

// PlayerServer .
type PlayerServer struct {
	store PlayerStore
	// router *http.ServeMux
	http.Handler
}

// NewPlayerServer return a new PlayerServer instance
// func NewPlayerServer(store PlayerStore) *PlayerServer {
// 	p := &PlayerServer{
// 		store:  store,
// 		router: http.NewServeMux(),
// 	}
// 	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// 	p.router.Handle("/players", http.HandlerFunc(p.playersHandler))
// 	return p
// }

// NewPlayerServer return a new PlayerServer instance
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{store: store}
	// p.Handler = p.register()

	mux := http.NewServeMux()
	mux.Handle("/league", http.HandlerFunc(p.leagueHandler))
	mux.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = mux

	return p
}

func (p *PlayerServer) register() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	return router
}

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// player := r.URL.Path[len("/players/"):]
// score := p.store.GetPlayerScore(player)
// if score == 0 {
// 	w.WriteHeader(http.StatusNotFound)
// }
// fmt.Fprint(w, score)

// router := http.NewServeMux()
// router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// router.Handle("/players", http.HandlerFunc(p.playersHandler))
// p.router.ServeHTTP(w, r)
// }

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	log.Printf("showScore() of player %q", player)
	// score := p.store.GetPlayerScore(player)
	// if score == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// }
	// fmt.Fprint(w, score)

	foundPlayer := p.store.GetLeague().Find(player)
	if foundPlayer == nil {
		p.JSON(w, http.StatusNotFound, map[string]any{
			"code": http.StatusNotFound,
			"msg":  fmt.Sprintf("player not found: %s", player),
		})
		return
	}
	p.JSON(w, http.StatusOK, foundPlayer)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	log.Printf("processWin() of player %q", player)
	p.store.RecordWinPlayer(player)
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w, http.StatusText(http.StatusOK))
	p.JSON(w, http.StatusOK, p.store.GetLeague().Find(player))
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request of leagueHandler processing ...")
	// w.Header().Set("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(p.store.GetLeague())
	p.JSON(w, http.StatusOK, p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request of playersHandler processing ...")
	player := r.URL.Path[len("/players/"):]
	log.Printf("playersHandler() get player %q", player)
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

// JSON write json fomat response to the player server
func (p *PlayerServer) JSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
