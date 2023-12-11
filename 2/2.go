package main

import (
	"fmt"
	"sync"
)

func sqrt(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	fmt.Println(number * number)
}

func main() {
	var wg sync.WaitGroup
	array := [5]int{2, 4, 6, 8, 10}

	for _, v := range array {
		wg.Add(1)
		go sqrt(&wg, v)
	}

	wg.Wait()
}
