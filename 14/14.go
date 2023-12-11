package main

import (
	"fmt"
	"reflect"
)

// We use reflection package in order to get types of our elements
// We use interface to be able to receive any type in our function
func checkType(variable interface{}) {
	fmt.Println(reflect.TypeOf(variable))
}

func main() {
	x := make(map[int]string)
	checkType(x)
}
