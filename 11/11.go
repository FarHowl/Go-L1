package main

import "fmt"

func intersection(set1 []int, set2 []int) {
	setMap := make(map[int]bool)

	for _, v := range set1 {
		setMap[v] = true
	}

	// There can only be unique values in the hashMap, so we use this feature
	for _, v := range set2 {
		if setMap[v] {
			fmt.Printf("%d ", v)
		}
	}
}

func main() {
	set1 := []int{1, 3, 2, 5, 4}
	set2 := []int{5, 8, 3, 9, 1}
	fmt.Println("Intersection :")
	intersection(set1, set2)
}
