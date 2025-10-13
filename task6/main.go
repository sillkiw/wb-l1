package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	byCondition()
	closeStopBroadcast()
	contextCancel()
	goexit()
}

// Обычное условие внутри горутины
func byCondition() {
	fmt.Println("\n-- 1) by condition --")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			if i == 3 {
				fmt.Println("stop by condition at i=3")
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	wg.Wait()
}

// Стоп через закрытие канала
func closeStopBroadcast() {
	fmt.Println("\n-- 2) close(stop) broadcast --")
	var wg sync.WaitGroup
	stop := make(chan struct{})

	worker := func(id int) {
		defer wg.Done()
		for {
			select {
			case <-stop:
				fmt.Println("worker", id, "stopped")
				return
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}

	wg.Add(2)
	go worker(1)
	go worker(2)

	time.Sleep(40 * time.Millisecond)
	close(stop)
	wg.Wait()
}

// Контекст: отмена WithCancel()
func contextCancel() {
	fmt.Println("\n-- 4) context cancel --")
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		fmt.Println("stopped via context:", ctx.Err())
	}()

	time.Sleep(40 * time.Millisecond)
	cancel()
	wg.Wait()
}

// runtime.Goexit()
func goexit() {
	fmt.Println("\n-- 5) runtime.Goexit() --")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("defer before Goexit executed")
			wg.Done()
		}()
		runtime.Goexit()
	}()
	wg.Wait()
}
