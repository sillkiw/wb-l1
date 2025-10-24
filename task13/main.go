package main

import "fmt"

func swapArfOp(a *int, b *int) {
	*a -= *b
	*b += *a
	*a = *b - *a
}

func swapXOR(a *int, b *int) {
	if a == b {
		return
	}
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	var (
		a int = -13
		b int = 13
	)

	swapArfOp(&a, &b)
	fmt.Println("a =", a, "b =", b)

	swapXOR(&a, &b)
	fmt.Println("a =", a, "b =", b)
}
