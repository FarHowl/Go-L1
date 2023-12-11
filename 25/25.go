package main

import (
	"fmt"
	"time"
)

// Simple for loop that checks current time and start time values
func sleep(sleepTime int64) {
	startTime := time.Now().UnixMilli()
	currentTime := time.Now().UnixMilli()

	for currentTime-startTime < sleepTime {
		currentTime = time.Now().UnixMilli()
	}
}

func main() {
	sleep(2000)
	fmt.Println("I am here")
}
