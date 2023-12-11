package main

import (
	"fmt"
	"strings"
)

func reverseString(originalStr *string) {
	var builder strings.Builder
	str := []rune(*originalStr)

	// Reverse loop
	for i := len(str) - 1; i >= 0; i-- {
		builder.WriteRune(str[i])
	}

	fmt.Println(builder.String())
}

func main() {
	str := "abcdefg"

	reverseString(&str)
}
