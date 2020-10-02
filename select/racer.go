package racer

import (
	"net/http"
	"time"
)

// Racer returns a URL the loads faster
func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	durationB := time.Since(startB)

	if durationA < durationB {
		winner = a
	} else {
		winner = b
	}

	return
}
