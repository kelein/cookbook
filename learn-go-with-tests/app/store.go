package app

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
