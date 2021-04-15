package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/brighteyed/learn-go-with-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Lets play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.Alerter))
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}