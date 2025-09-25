package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var possible = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	amwork := flag.Int("amount", 5, "amount of workers")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ch := make(chan string, *amwork)
	var wg sync.WaitGroup

	for i := 0; i < *amwork; i++ {
		wg.Add(1)
		go func(id int) {
			worker(id, ch)
			wg.Done()
		}(i)
	}

	produce(ctx, ch)

	wg.Wait()
	fmt.Println("graceful shutdown complete")

}

func produce(ctx context.Context, ch chan<- string) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- randString(5):
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func randString(length int) string {
	res := make([]rune, length)
	for i := range res {
		res[i] = possible[rand.Intn(len(possible))]
	}
	return string(res)
}

func worker(id int, ch chan string) {
	for val := range ch {
		fmt.Printf("worker %d print from chanel: %s\n", id, val)
		time.Sleep(time.Second)
	}

}
