package main

import (
	"fmt"
	"unicode"
)

// We use the feature of hashMap (unique keys) to store our runes until there is a similar one
func allSymbolsUnique(str string) bool {
	set := make(map[rune]bool)
	for _, r := range str {
		lowerRune := unicode.ToLower(r)
		if !set[lowerRune] {
			set[lowerRune] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	str := "abCdefAaf"
	fmt.Println(allSymbolsUnique(str))
}
