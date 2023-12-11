package main

import "fmt"

func main() {
	// Arithmetic method to swap numbers
	myArray := []int{1, 2}
	myArray[0] = myArray[0] + myArray[1]
	myArray[1] = myArray[0] - myArray[1]
	myArray[0] = myArray[0] - myArray[1]

	fmt.Println(myArray)
}
