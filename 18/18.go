package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	counter int64
}

// We use atomics as they allow us to use a few goroutines which don`t block each other
func increment(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter.counter, 1)
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.counter)
}
