package app

import (
	"io"
)

// FileSystemStore with file system
type FileSystemStore struct {
	// database io.Reader
	database io.ReadSeeker
}

// GetLeague return player's league from FileSystemStore
// func (f *FileSystemStore) GetLeague() []Player {
// 	league := make([]Player, 0)
// 	json.NewDecoder(f.database).Decode(&league)
// 	return league
// }

// GetLeague return player's league from FileSystemStore
func (f *FileSystemStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore get player score by name
func (f *FileSystemStore) GetPlayerScore(name string) int {
	for _, player := range f.GetLeague() {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}

// RecordWinPlayer records player who has won
func (f *FileSystemStore) RecordWinPlayer(name string) {}
