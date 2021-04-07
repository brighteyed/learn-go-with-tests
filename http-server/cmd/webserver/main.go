package main

import (
	"log"
	"net/http"

	poker "github.com/brighteyed/learn-go-with-tests/http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	if err := http.ListenAndServe(":5000", poker.NewPlayerServer(store)); err != nil {
		log.Fatalf("could not listen on port 5000, %v", err)
	}
}
