package main

import (
	"fmt"
	"strings"
)

// Same algorithm as in 19th exercise
func reverseWords(originalString *string) {
	words := strings.Fields(*originalString)
	var builder strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i] + " ")
	}

	*originalString = builder.String()
	*originalString = strings.Trim((*originalString), " ")
}

func main() {
	str := "cat dog snow"
	reverseWords(&str)

	fmt.Println(str)
}
