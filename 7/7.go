package main

import (
	"fmt"
	"sync"
)

// We use mutex to avoid race condition
func writeToMap(myMap *map[int]string, number int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	(*myMap)[number] = "a"
	mu.Unlock()
}

func main() {
	myMap := make(map[int]string)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeToMap(&myMap, i, &mu, &wg)
	}

	for i, v := range myMap {
		fmt.Println(i, v)
	}

	wg.Wait()
}
