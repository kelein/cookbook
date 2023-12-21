package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelein/cookbook/learn-go-with-tests/app"
)

const (
	dbFileName = "game.db"
	serverPort = 5000
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
}

func main() {
	// db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatalf("open db file %q err: %v", dbFileName, err)
	// }

	store, err := app.NewFileSystemPlayerStoreLocal(dbFileName)
	if err != nil {
		log.Fatalf("init store err: %v", err)
	}

	server := app.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), server))
}
