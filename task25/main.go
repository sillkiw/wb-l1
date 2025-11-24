package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Sleep 2 seconds")
	sleep(2 * time.Second)
	fmt.Println("Done")
}

func sleep(d time.Duration) {
	if d <= 0 {
		return
	}
	t := time.NewTimer(d)
	<-t.C
}
