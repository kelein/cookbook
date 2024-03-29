package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kelein/cookbook/learn-go-with-tests/app"
)

// PlayerClient client for Player
type PlayerClient struct {
	playerStore app.PlayerStore
	input       *bufio.Scanner
}

// NewPlayerClient create a PlayerClient instance
func NewPlayerClient(store app.PlayerStore, input io.Reader) *PlayerClient {
	return &PlayerClient{
		playerStore: store,
		input:       bufio.NewScanner(input),
	}
}

// PlayPoker of PlayerClient
func (pc *PlayerClient) PlayPoker() {
	content := pc.readLine()
	winner := pc.parseWinner(content)
	pc.playerStore.RecordWinPlayer(winner)
}

func (pc *PlayerClient) readLine() string {
	pc.input.Scan()
	return pc.input.Text()
}

func (pc *PlayerClient) parseWinner(text string) string {
	return strings.Replace(text, " wins", "", 1)
}

const dbFileName = "game.json"

func main() {
	// db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatalf("open store file failed: %v", err)
	// }

	store, err := app.NewFileSystemPlayerStoreLocal(dbFileName)
	if err != nil {
		log.Fatalf("init store failed: %v", err)
	}

	fmt.Println("-------- Let's play poker --------")
	fmt.Println("Type {Name} wins to record")
	cli := NewPlayerClient(store, os.Stdin)
	cli.PlayPoker()
}
