package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	var cnt Counter

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(cnt *Counter, i int) {
			defer wg.Done()
			for j := 0; j < i+1; j++ {
				cnt.Increment()
				fmt.Printf("Increment by %d goroutine\n", i)
				fmt.Printf("Get value %d by %d goroutine\n", cnt.Get(), i)
			}
		}(&cnt, i)
	}

	wg.Wait()
	fmt.Println("Final value:", cnt.Get())
}
