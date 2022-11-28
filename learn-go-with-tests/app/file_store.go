package app

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore with file system
type FileSystemPlayerStore struct {
	// database io.Reader
	database io.ReadWriteSeeker
	league   League
}

// NewFileSystemPlayerStore create a FileSystemPlayerStore instance
func NewFileSystemPlayerStore(db io.ReadWriteSeeker) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)
	return &FileSystemPlayerStore{
		database: db,
		league:   league,
	}
}

// GetLeague return player's league from FileSystemStore
// func (f *FileSystemStore) GetLeague() []Player {
// 	league := make([]Player, 0)
// 	json.NewDecoder(f.database).Decode(&league)
// 	return league
// }

// GetLeague return player's league from FileSystemStore
// func (f *FileSystemPlayerStore) GetLeague() League {
// 	f.database.Seek(0, 0)
// 	league, _ := NewLeague(f.database)
// 	return league
// }

// GetLeague return player's league from FileSystemStore
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

// GetPlayerScore get player score by name
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	// for _, player := range f.GetLeague() {
	// 	if player.Name == name {
	// 		return player.Wins
	// 	}
	// }
	// return 0

	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWinPlayer records player who has won
// func (f *FileSystemPlayerStore) RecordWinPlayer(name string) {
// 	league := f.GetLeague()
// 	for i, player := range league {
// 		if player.Name == name {
// 			league[i].Wins++
// 		}
// 	}
// 	f.database.Seek(0, 0)
// 	json.NewEncoder(f.database).Encode(league)
// }

// RecordWinPlayer records player who has won
func (f *FileSystemPlayerStore) RecordWinPlayer(name string) {
	player := f.league.Find(name)
	if player == nil {
		newPlayer := Player{Name: name, Wins: 1}
		f.league = append(f.league, newPlayer)
	} else {
		player.Wins++
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
