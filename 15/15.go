package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var justString string

// We create random huge string
func createHugeString(size int) string {
	var builder strings.Builder
	for i := 0; i < size; i++ {
		builder.WriteRune(rune(rand.Intn(26) + 'a'))
	}

	return builder.String()
}

func someFunc() {
	v := createHugeString(1 << 10)

	// Then we use strings.Builder in order to copy first 100 runes into our justString variable. No references
	var builder strings.Builder
	for i, rune := range v {
		if i < 100 {
			builder.WriteRune(rune)
		} else {
			justString = builder.String()
			break
		}
	}

	fmt.Println(justString)
}

func main() {
	someFunc()
}
