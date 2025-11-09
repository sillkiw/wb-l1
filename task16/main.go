package main

import "fmt"

func main() {
	test := []int{2, 3, 4, 5, 1, 2, 3, 4}
	fmt.Println(QuickSort(test))
}

func QuickSort(a []int) []int {
	return quickSortR(a, 0, len(a)-1)
}

func quickSortR(a []int, low, high int) []int {
	if low < high {
		var pivot int
		a, pivot = partition(a, low, high)
		a = quickSortR(a, low, pivot-1)
		a = quickSortR(a, pivot+1, high)
	}
	return a
}

func partition(a []int, low, high int) ([]int, int) {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return a, i
}
