package main

import (
	"fmt"
)

func groupTemp(temp []float64) map[int][]float64 {

	res := make(map[int][]float64)

	for _, t := range temp {
		st := int(t/10) * 10
		res[st] = append(res[st], t)
	}
	return res
}

func main() {
	res := groupTemp([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5})
	fmt.Println(res)
}
