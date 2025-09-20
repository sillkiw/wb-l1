package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var possible = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {

	amwork := flag.Int("amount", 5, "amount of workers")
	flag.Parse()

	ch := make(chan string)

	for i := 0; i < *amwork; i++ {
		go worker(i, ch)
	}

	for {
		ch <- randString(5)
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
