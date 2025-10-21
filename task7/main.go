package main

import (
	"fmt"
	"sync"
)

var a map[rune]int
var wg sync.WaitGroup

func work(number int, mutex *sync.Mutex) {

	mutex.Lock()

	a = make(map[rune]int)

	word := "absdfasdfxcvzxasdfaszxcvasdfzxcvzxcv"
	for _, k := range word {
		if _, seen := a[k]; !seen {
			a[k] = 0
		} else {
			a[k]++
		}
	}
	fmt.Println("Goroutine", number, "-", a)
	mutex.Unlock()

	wg.Done()
}

func main() {

	var mutex sync.Mutex

	goroutines_count := 4

	wg.Add(goroutines_count)

	for i := 1; i <= goroutines_count; i++ {
		go work(i, &mutex)
	}
	wg.Wait()

	fmt.Println("The End")
}
