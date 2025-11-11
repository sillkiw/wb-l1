package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	fmt.Println(reverseWords(line))

}

func reverseWords(s string) string {
	res := []rune(s)
	if len(res) == 0 {
		return s
	}
	reverse(res, 0, len(res)-1)

	start := 0
	for i := 0; i <= len(res); i++ {
		if i == len(res) || unicode.IsSpace(res[i]) {
			reverse(res, start, i-1)
			start = i + 1
		}
	}
	return string(res)
}

func reverse(s []rune, l, r int) {
	for i, j := l, r; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
