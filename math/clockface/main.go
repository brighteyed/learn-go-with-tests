package main

import (
	"os"
	"time"

	"github.com/brighteyed/learn-go-with-tests/math/svg"
)

func main() {
	svg.Write(os.Stdout, time.Now())
}
