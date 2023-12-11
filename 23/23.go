package main

import (
	"errors"
	"fmt"
)

// We create newSlice, copying certain elements from the old one except the one we want to delete
func deleteElement(slice *[]int, index int) error {
	if index < len(*slice) && index >= 0 {
		var newSlice []int
		if index == 0 {
			newSlice = append(newSlice, (*slice)[1:]...)
		} else if index == len(*slice)-1 {
			newSlice = append(newSlice, (*slice)[:index]...)
		} else {
			newSlice = append(newSlice, (*slice)[:index]...)
			newSlice = append(newSlice, (*slice)[index+1:]...)
		}

		(*slice) = newSlice

		return nil
	} else {
		return errors.New("error: index out of range")
	}
}

func main() {
	slice := []int{1, 2, 3, 4}

	err := deleteElement(&slice, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slice)
	}
}
