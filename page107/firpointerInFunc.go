package main

import (
	"fmt"
)

func swap(fir *int, sec *int) {
	temp := 0
	temp = *fir
	*fir = *sec
	*sec = temp
	fmt.Println(*fir, *sec)
}

func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)
	swap(&a, &b) //&
	fmt.Println(a, b)
	//fmt.Printf("%.2f\n"), math Sqrt(-25)
}
