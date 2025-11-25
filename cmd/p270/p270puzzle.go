package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println("S.C. pop.:", population)
	population += 1
	fmt.Println("pop. after immigrat.:", population)
}
