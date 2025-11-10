package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	//GetFloats(): 위 깃 링크(패키지)에서 정의됨.
	weights, err := datafile.GetFloats("wlist.txt")
	//error 복습 좀..................
	if err != nil {
		log.Fatal(err)
	}
	//var sum float64 = 0
	sum := 0.0
	//for _, num := range nums {
	for i := 0; i < len(weights); i++ {
		sum += weights[i]
	}
	counting := float64(len(weights))
	fmt.Printf("Average: %0.2f\n", sum/counting)
}
