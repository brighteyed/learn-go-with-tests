package main

import (
	"log"
	"net/http"

	poker "github.com/brighteyed/learn-go-with-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.Alerter))
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("could not create server, %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000, %v", err)
	}
}
