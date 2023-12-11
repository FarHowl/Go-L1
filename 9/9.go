package main

import (
	"fmt"
	"sync"
)

// Goroutine which sends numbers to first channel
func addNumbers(number int, arrayChan chan int) {
	arrayChan <- number
}

// Goroutine which receives numbers from the first channel and sends powed numbers to seconds channel
func powNumbers(arrayChan chan int, powChan chan int) {
	for v := range arrayChan {
		powChan <- v * v
	}
}

// Goroutine which receives numbers from the second channel and prints them
func printNumbers(powChan chan int, wg *sync.WaitGroup) {
	for v := range powChan {
		fmt.Println(v)
		wg.Done()
	}
}

func main() {
	var arrayChan chan int = make(chan int)
	var powChannel chan int = make(chan int)

	var wg sync.WaitGroup

	array := []int{1, 2, 3, 4, 5}

	go powNumbers(arrayChan, powChannel)
	go printNumbers(powChannel, &wg)

	for _, v := range array {
		wg.Add(1)
		go addNumbers(v, arrayChan)
	}

	wg.Wait()
}
