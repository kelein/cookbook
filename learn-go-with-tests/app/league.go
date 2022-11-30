package app

import (
	"encoding/json"
	"io"
)

// League of Players
type League []*Player

// Find found player by name
func (l League) Find(name string) *Player {
	for _, p := range l {
		if p.Name == name {
			return p
		}
	}
	return nil
}

// NewLeague create a player league
func NewLeague(r io.Reader) (League, error) {
	league := make([]*Player, 0)
	err := json.NewDecoder(r).Decode(&league)
	return league, err
}
