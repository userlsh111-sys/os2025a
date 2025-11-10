package main

import "fmt"

func main() {
	//ERROR: var presidents []string{"", "", "yoon", "lee"}
	presidents := []string{"", "", "yoon", "lee"}

	//var presidents make([]string, 4)
	//presidents[0] = "park", presidents[1] = "moon"
	//이렇게도 ㄱㄴ

	for _, president := range presidents {
		fmt.Println(president)
	}
}
