package main

import (
	"fmt"
)

func main() {
	c1 := &digitalClock{src: format24{23, 5}}
	c1.ShowTime()

	// Применение адаптера для digitalClock
	twelveAm := format12{1, 35, "AM"}
	clc := &digitalClock{src: format12Adapter{twelveAm}}
	clc.ShowTime()
}

type time24Format interface {
	Extract() string
}

// digitalClock работает только с time24Format
type digitalClock struct {
	src time24Format
}

func (c *digitalClock) ShowTime() {
	fmt.Println(c.src.Extract())
}

// format24 - структура удовлет. интерфейсу time24Format
type format24 struct {
	Hour int
	Min  int
}

func (f format24) Extract() string {
	return fmt.Sprintf("%02d:%02d", f.Hour, f.Min)
}

// format12 - структура не удовлт. интерфейсу
type format12 struct {
	Hour int
	Min  int
	AmPm string
}

// format12Adapter адаптер для 12-часового формата
type format12Adapter struct {
	f format12
}

// Extract() переводит 12 часовой формат в 24 часовой
func (a format12Adapter) Extract() string {
	ampm := a.f.AmPm
	h := a.f.Hour % 12
	if ampm == "PM" {
		h += 12
	}
	return fmt.Sprintf("%02d:%02d", h, a.f.Min)

}
