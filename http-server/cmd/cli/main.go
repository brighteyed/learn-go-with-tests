package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/brighteyed/learn-go-with-tests/http-server"
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

	cli := poker.NewCLI(store, os.Stdin)
	cli.PlayPoker()
}
