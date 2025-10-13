package main

import (
	"fmt"
	"reflect"
)

func main() { //main 필수
	var length float64 = 3.2
	var width int = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))

	fmt.Println("형변환(정수로.): ", reflect.TypeOf(int(length)))
	fmt.Println("원본: ", reflect.TypeOf(length))
	//TypeOf: 말 그대로 자료형
}
