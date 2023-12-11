package main

import "fmt"

func main() {
	myString := []string{"cat", "cat", "dog", "cat", "tree"}

	setMap := make(map[string]bool)
	set := []string{}

	// There can only be unique values in the hashMap, so we use this feature
	for _, v := range myString {
		if !setMap[v] {
			setMap[v] = true
			set = append(set, v)
		}
	}

	fmt.Println(set)
}
