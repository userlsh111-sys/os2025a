package main

import "fmt"

func main() {
	var width, height, area float64

	width = 2.525
	height = 5.555
	area = width * height
	fmt.Printf("The area is %.2f.\n", area) //C에서처럼 %와 ,로.

	width = 10.101
	height = 2.525
	area = width * height
	fmt.Printf("The area is %.2f.\n", area)

	fmt.Println("====Prinf로 줄어든 자릿수====")
}
