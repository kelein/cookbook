package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cookbook/learn-go-with-tests/app"
)

const dbFile = "game.db"

var port = 5000

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
}

func main() {
	db, err := os.OpenFile(dbFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open db file %q err: %v", dbFile, err)
	}

	store := app.NewFileSystemPlayerStore(db)
	server := app.NewPlayerServer(store)
	http.ListenAndServe(fmt.Sprintf(":%d", port), server)
}
