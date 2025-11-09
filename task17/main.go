package main

import (
	"fmt"
)

func main() {
	a := []int{-1000000, -5000, 0, 1000000000, 1000000001}
	want := 0

	fmt.Println("recursive:", binSearchR(a, want))
	fmt.Println("for", binSearchI(a, want))

}

// binSearchR - рекурсивная реализация бин поиска
func binSearchR(a []int, want int) int {

	var binSearchRec func([]int, int, int, int) int
	binSearchRec = func(a []int, l, r, want int) int {
		if l > r {
			return -1
		}
		mid := l + (r-l)/2
		if a[mid] > want {
			return binSearchRec(a, l, mid-1, want)
		} else if a[mid] < want {
			return binSearchRec(a, mid+1, r, want)
		}
		return mid
	}
	return binSearchRec(a, 0, len(a)-1, want)
}

// binSearchI - реализцая бин поиска через цикл
func binSearchI(a []int, want int) int {
	l, r := 0, len(a)

	for l < r {
		mid := l + (r-l)/2

		if a[mid] > want {
			r = mid - 1
		} else if a[mid] < want {
			l = mid + 1
		} else {
			return mid
		}

	}
	return -1
}
