package main

import (
	"fmt"
	"reflect"
)

func defineType(src interface{}) {
	switch src.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.TypeOf(src).Kind() == reflect.Chan {
			fmt.Println("chan")
		}
	}

}

func main() {
	defineType(8)
	defineType("hi")
	defineType(true)
	defineType(make(chan int))

}
