package main

import "fmt"

func intersection(set1 []int, set2 []int) []int {

	match := make(map[int]struct{}, len(set1))
	for _, v := range set1 {
		match[v] = struct{}{}
	}

	interSet := make(map[int]struct{})
	for _, v := range set2 {
		if _, seen := match[v]; seen {
			interSet[v] = struct{}{}
		}
	}

	ans := make([]int, 0, len(interSet))
	for v := range interSet {
		ans = append(ans, v)
	}
	return ans
}

func main() {

	set1 := []int{1, 2, 3}
	set2 := []int{2, 3, 4, 3}

	fmt.Println(intersection(set1, set2))

}
