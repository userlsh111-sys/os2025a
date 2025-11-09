package main

import (
	"fmt"
)

func main() {
	numbers := [4]int{19, -72, 12, -21}

	for i, number := range numbers {
		fmt.Println(i, number)
	}
}
