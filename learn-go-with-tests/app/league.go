package app

import (
	"encoding/json"
	"io"
)

// NewLeague create a player league
func NewLeague(r io.Reader) ([]Player, error) {
	league := make([]Player, 0)
	err := json.NewDecoder(r).Decode(&league)
	return league, err
}
