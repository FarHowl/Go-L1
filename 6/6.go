package main

import (
	"context"
	"fmt"
)

func worker_context(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Worker was stopped by context")
}

func worker_channel(stop chan bool) {
	select {
	case <-stop:
		fmt.Println("Worker was stopped by channel")
	default:
		fmt.Println("Working")
	}
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

	fmt.Println("Goroutines were stopped")
}
