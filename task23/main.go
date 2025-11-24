package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int = 4

	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]
	fmt.Println(a)
}
