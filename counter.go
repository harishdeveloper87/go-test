package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	for v := 0; v < 10; v++ {
		go func(v int) {
			doublev := callDouble(v)
			log.Printf("Thread %d returned: %d", v, doublev)
		}(v)
	}

	time.Sleep(time.Second * 10)
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu    *sync.Mutex
	count int64
}

var (
	counterLimit = 5          // limit when we want to stop go routine.
	sc           = NewCount() // initiate the global counter.
)
// NewCount returns the pointer instance of SafeCounter.
func NewCount() *SafeCounter {
	return &SafeCounter{mu: new(sync.Mutex), count: 0}
}

// Inc increments the counter.
func (c *SafeCounter) Inc() {
	// Lock so only one goroutine at a time can access the c.count.
	c.mu.Lock()

	// increment counter by one.
	c.count++

	// Unlock will unlock the mutex immediately prior to Inc returning.
	c.mu.Unlock()
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int64 {
	// Lock so only one goroutine at a time can access the c.count.
	c.mu.Lock()

	// Unlock will unlock the mutex immediately prior to Value returning.
	// defer is extremely useful because without it,
	// you would have to remember to unlock the mutex in every place that the function could possibly return
	defer c.mu.Unlock()

	return c.count
}
func callDouble(v int) int {
	// increase the global counter by one.
	sc.Inc()

	// stop returning value from the double(v) when counter reaches to limit.
	if sc.Value() > int64(counterLimit) {
		return v
	}
	return double(v)
}

func double(v int) int {
	time.Sleep(time.Second)
	return v * 2
}
