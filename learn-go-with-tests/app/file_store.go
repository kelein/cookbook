package app

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type tape struct {
	// file io.ReadWriteSeeker
	file *os.File
}

func (t *tape) Write(p []byte) (int, error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

func (t *tape) Seek(offset int64, whence int) (int64, error) {
	return t.file.Seek(offset, whence)
}

// FileSystemPlayerStore with file system
type FileSystemPlayerStore struct {
	// database io.Reader
	// database io.Writer
	// database io.ReadWriteSeeker

	database *tape
	league   League
}

// NewFileSystemPlayerStore create a FileSystemPlayerStore instance
// func NewFileSystemPlayerStore(db io.ReadWriteSeeker) *FileSystemPlayerStore {
// 	db.Seek(0, 0)
// 	league, _ := NewLeague(db)
// 	return &FileSystemPlayerStore{
// 		// database: db,
// 		database: &tape{db},
// 		league:   league,
// 	}
// }

// NewFileSystemPlayerStore create a FileSystemPlayerStore instance
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	if err := initStoreFile(file); err != nil {
		return nil, errors.Wrap(err, "init store file failed")
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, errors.Wrap(err, "load league from file failed")
	}

	return &FileSystemPlayerStore{
		database: &tape{file},
		league:   league,
	}, nil
}

func initStoreFile(f *os.File) error {
	f.Seek(0, 0)
	info, err := f.Stat()
	if err != nil {
		return errors.Wrap(err, "get file info failed")
	}
	if info.Size() == 0 {
		f.Write([]byte("[]"))
		f.Seek(0, 0)
	}
	return nil
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
