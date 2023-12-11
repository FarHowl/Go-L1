package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var myMap map[int][]float32 = make(map[int][]float32)
	array := []float32{-25.4, -27.0, 13.0, 10.0, 15.5, 24.5, -21.0, 32.5}

	// We are sorting our array in order to find left border of our sets
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	currentSet := int(array[0]) / 10 * 10
	myMap[currentSet] = []float32{array[0]}

	for i := 1; i < len(array); i++ {
		// If the difference between set value and the element is higher than 10
		if math.Abs(float64(currentSet)-float64(array[i])) <= 10 {
			// Add new element to our existing slice
			myMap[currentSet] = append(myMap[currentSet], array[i])
		} else {
			// Else create new set
			currentSet = int(array[i]) / 10 * 10
			myMap[currentSet] = []float32{array[i]}
		}
	}

	fmt.Println(myMap)
}
