package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var name float32
	var name float64

	//name = "InhaKim"

	//var name = "InhaKim"
	//name := "InhaKim"

	name = 2.71
	fmt.Println(name, reflect.TypeOf(name))
}
