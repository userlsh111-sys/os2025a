package main

import (
	"fmt"
)

func main() {
	var width, height, area float64

	width = 2.525
	height = 5.555
	area = width * height
	fmt.Println("The area is", area)

	width = 10.101 //변수 다시 선언하면..
	height = 2.525
	area = width * height //관련 변수도 다시 선언해야.
	fmt.Println("The area is", area)

	fmt.Println("====꽤 많은 자릿수====")
}
