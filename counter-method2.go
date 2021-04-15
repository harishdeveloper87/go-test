package main

import (
	"log"
	"time"
	"github.com/zenthangplus/goccm"
)

func main() {
    c := goccm.New(5)
	for v := 0; v < 10; v++ {
	  c.Wait()
		go func(v int) {
			doublev := callDouble(v)
			log.Printf("Thread %d returned: %d", v, doublev)
			c.Done()
		}(v)
	}
c.WaitAllDone()
	time.Sleep(time.Second * 10)
}

func callDouble(v int) int {
	return double(v)
}

func double(v int) int {
	time.Sleep(time.Second)
	return v * 2
}
