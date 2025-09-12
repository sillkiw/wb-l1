package main

import "fmt"

type Human struct {
	name   string
	age    int8
	height int16
	weight int16
	sex    rune
}

func (h Human) GetName() string {
	return h.name
}

type Action struct {
	Human
}

func main() {
	human := Human{
		name:   "Alex",
		age:    40,
		height: 180,
		weight: 80,
		sex:    'M',
	}
	action := Action{Human: human}
	fmt.Println(action.GetName())
}
