package main

import (
	"fmt"
	"math/big"
)

func main() {

	a, b := new(big.Int), new(big.Int)
	a.SetString("240000000000000000000000000", 10)
	b.SetString("24000000000000000000000000", 10)

	fmt.Println("a + b =", new(big.Int).Add(a, b))
	fmt.Println("a - b =", new(big.Int).Sub(a, b))
	fmt.Println("a * b =", new(big.Int).Mul(a, b))
	fmt.Println("a / b =", new(big.Int).Div(a, b))

}
