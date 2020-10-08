package counter

import "sync"

// Counter is a threadsafe counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments a counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns a current value
func (c *Counter) Value() int {
	return c.value
}

// NewCounter is a constructor for counter
func NewCounter() *Counter {
	return &Counter{}
}
