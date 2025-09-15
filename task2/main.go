package main

import (
	"fmt"
	"time"
)

func evSquare(a []int, i int) {
	for _, el := range a {
		fmt.Printf("%d^2 = %d, выведено с горутины %d\n", el, el*el, i)
		time.Sleep(time.Microsecond)
	}
}

func main() {
	var a = []int{2, 4, 6, 8, 10}

	for i := 1; i <= 3; i++ {
		go evSquare(a, i)
	}
	time.Sleep(1 * time.Second)
}
