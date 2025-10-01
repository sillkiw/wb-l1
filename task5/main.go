package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var possible = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	t := flag.Int("dur", 1, "duration")
	flag.Parse()

	ch := make(chan string)
	go func(ch chan string) {
		for msg := range ch {
			fmt.Println(msg)
		}
	}(ch)

	produce(ch, *t)

}

func produce(ch chan string, t int) {
	defer close(ch)
	timeout := time.After(time.Duration(t) * time.Second)
	for {

		select {
		case ch <- randString(10):
			time.Sleep(10 * time.Millisecond)
		case <-timeout:
			return
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
