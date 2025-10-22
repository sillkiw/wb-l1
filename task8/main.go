package main

import (
	"fmt"
)

func main() {

	for {
		var a int64
		fmt.Println("Введите число")
		fmt.Scan(&a)

		var i int8
		fmt.Println("Введите бит числа, который хотите поменять")
		fmt.Scan(&i)

		var ch rune
		fmt.Printf("Установить %d бит в 0 или 1 (введите 0 или 1)\n", i)
		fmt.Scan(&ch)

		var mask int64 = 1 << (i - 1)
		if ch == 1 {
			a |= mask
		} else {
			a &= (^mask)
		}
		fmt.Println(a)
	}
}
