package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Writing data to channel
func writeData(ch chan int) {
	for {
		ch <- rand.Intn(100)
		time.Sleep(time.Second)
	}
}

// Reading data from channel in a for loop until the channel is closed
func readData(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Worker exited")
}

func main() {
	var N int
	fmt.Print("Введите количество воркеров N: ")
	fmt.Scanln(&N)

	var ch chan int = make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go readData(&wg, ch)
	}

	// Make a channel for one interrupting signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)

	// Run goroutine, which will close the main channel, when the interrupting signal is sent
	go func() {
		<-sigCh
		close(ch)
	}()

	// Run writing data in a goroutine in order to step forward to listening for WaitGroup
	go writeData(ch)

	wg.Wait()
	fmt.Println("All workers have exited")
}
