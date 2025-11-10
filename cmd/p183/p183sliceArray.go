package main

import "fmt"

func main() {
	presidents := []string{"park", "moon", "yoon", "lee"}

	somePres := presidents[:3] //presidents[0]="park"
	somePres[0] = "sejong"

	for _, president := range presidents {
		fmt.Println(president)
	}
	fmt.Println("========아래: some========")
	for i := 0; i < len(somePres); i++ {
		fmt.Println(somePres[i])
	}
}
