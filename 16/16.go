package main

import (
	"fmt"
	"math/rand"
)

func swap(i, j int, array *[]int) {
	temp := (*array)[i]
	(*array)[i] = (*array)[j]
	(*array)[j] = temp
}

// This quicksort function works on the method of two pointers. Left and right pointers go from opposite sides and swap elements if it`s needed until they meet each other. Then we do the same thing over left and right sides
func quicksort(lowIndex, highIndex int, array *[]int) {
	if lowIndex >= highIndex {
		return
	}

	// We want to use our last element as pivot, but we also want some random thing (Due to optimization. Choose random element and swap it with the last
	pivotIndex := rand.Intn(int(highIndex-lowIndex)) + lowIndex
	pivot := (*array)[pivotIndex]
	swap(pivotIndex, highIndex, array)

	lp := lowIndex
	rp := highIndex

	for lp < rp {
		for lp < rp && (*array)[lp] <= pivot {
			lp++
		}

		for lp < rp && (*array)[rp] >= pivot {
			rp--
		}

		swap(lp, rp, array)
	}

	// Swap leftPointer value with the last one, so all the values on the left side of this element are lower and all the values on the right side are higher
	if (*array)[lp] > pivot {
		swap(lp, highIndex, array)
	}

	// Do the whole thing to the left and right sides
	quicksort(lowIndex, lp-1, array)
	quicksort(lp+1, highIndex, array)
}

func main() {
	array := []int{}
	for i := 0; i < 100; i++ {
		array = append(array, rand.Intn(100))
	}

	quicksort(0, len(array)-1, &array)
	fmt.Println(array)
}
