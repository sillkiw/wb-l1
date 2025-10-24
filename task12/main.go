package main

import "fmt"

func ownSubset(src []string) []string {
	seen := make(map[string]struct{}, len(src))
	res := make([]string, 0, len(src))
	for _, v := range src {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		res = append(res, v)
	}
	return res
}

func main() {
	ts := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(ownSubset(ts))

}
