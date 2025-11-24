package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(unicSym(s))
}

func unicSym(s string) bool {
	set := map[rune]bool{}

	for _, l := range s {
		ch := unicode.ToLower(l)
		if set[ch] {
			return false
		}
		set[ch] = true
	}
	return true
}
