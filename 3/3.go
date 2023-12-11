package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// We use Atomics in order to avoid useless blocking for every goroutine
func sqrt(value int64, sum *int64, wg *sync.WaitGroup) {
	atomic.AddInt64(sum, value*value)
	wg.Done()
}

func main() {
	var sum int64
	array := [5]int64{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(array))

	for _, v := range array {
		go sqrt(v, &sum, &wg)
	}

	wg.Wait()
	fmt.Println(sum)
}
