package main

import (
	"context"
	"fmt"
	"sync"
)

func worker_context(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Worker was stopped by context")
}

func worker_channel(stop chan bool) {
	<-stop
	fmt.Println("Worker was stopped by channel")
}

func worker_wait_group(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine was stopped by wait group")
}

func main() {
	// Context
	ctx, cancel := context.WithCancel(context.Background())

	go worker_context(ctx)
	cancel()

	// Channel
	stop := make(chan bool)
	go worker_channel(stop)

	stop <- true

	// Wait Group
	var wg sync.WaitGroup
	wg.Add(1)
	go worker_wait_group(&wg)
	wg.Wait()

	fmt.Println("Goroutines were stopped")
}
