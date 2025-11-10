package main

import "fmt"

func main() {
	presidents := []string{"park", "moon", "yoon", "lee"}
	somePres := presidents[:3]
	somePres[0] = "sejong"

	for i := 0; i < len(somePres); i++ {
		fmt.Println(somePres[i])
	}
	fmt.Println("========아래: append 후========")

	somePres = append(somePres, "leesunggye", "jungjo", "kimku")

	for i := 0; i < len(somePres); i++ {
		fmt.Println(somePres[i])
	}
}
