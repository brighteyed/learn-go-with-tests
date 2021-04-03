package main

import (
	"log"
	"net/http"
)

func main() {
	println("http://localhost:5000")

	log.Fatal(http.ListenAndServe(":5000", NewPlayerServer(NewInMemoryPlayerStore())))
}
