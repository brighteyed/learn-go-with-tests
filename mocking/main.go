package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper can sleep for a while
type Sleeper interface {
	Sleep()
}

const countdownStart = 3
const finalWord = "Go!"

// Countdown 3 2 1 Go!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(out, finalWord)
}

// ConfigurableSleeper calls callback on each Sleep method call
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep calls a stored callback to sleep
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
