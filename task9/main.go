package main

import "fmt"

func generator(out chan<- int) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	for _, el := range a {
		out <- el
	}
	close(out)
}

func doubler(in <-chan int, out chan<- int) {
	for el := range in {
		out <- el * 2
	}
	close(out)
}

func main() {
	firstCh := make(chan int)
	secondCh := make(chan int)

	go generator(firstCh)
	go doubler(firstCh, secondCh)

	for el := range secondCh {
		fmt.Println(el)
	}

}
