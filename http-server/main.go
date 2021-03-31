package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	return 123, nil
}

func main() {
	println("http://localhost:5000")

	log.Fatal(http.ListenAndServe(":5000", &PlayerServer{
		store: &InMemoryPlayerStore{},
	}))
}
