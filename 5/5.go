package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func writeData(ch chan int) {
	for {
		ch <- rand.Intn(100)
		time.Sleep(time.Second)
	}
}

func readData(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	var N int
	fmt.Println("Введите время работы программы: ")
	fmt.Scanln(&N)

	var ch chan int = make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go writeData(ch)
	go readData(&wg, ch)

	go func() {
		time.Sleep(time.Second * time.Duration(N))
		close(ch)
	}()

	wg.Wait()
}
