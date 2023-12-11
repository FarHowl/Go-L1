package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// Simple algorithm. It uses sorted array and checks, if our value is higher or lower than the pivot (the middle value of the array). Do the same thing for the left and right sides
func binarySearch(array []int, value int) bool {
	if len(array) == 0 {
		return false
	}

	pivotIndex := len(array) / 2
	pivot := array[pivotIndex]

	if value == pivot {
		return true
	} else if value < pivot {
		return binarySearch(array[0:pivotIndex], value)
	} else {
		return binarySearch(array[pivotIndex+1:], value)
	}
}

func main() {
	var array []int
	for i := 0; i < 1000; i++ {
		array = append(array, rand.Intn(1000))
	}

	slices.Sort(array)

	fmt.Println(array)

	fmt.Println(binarySearch(array, 2))
}
