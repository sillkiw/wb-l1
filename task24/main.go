package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := NewPoint(4, 5)
	p2 := NewPoint(6, 7)
	fmt.Println("dist:", p1.Distance(p2))
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}
