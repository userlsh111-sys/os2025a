package main

import "fmt"

type meters float64
type kilometers float64
type miles float64

func (l meters) Tomiles() miles {
	return miles(l * 0.264)
}
func (m kilometers) Tomiles() miles {
	return miles(m * 0.000264)
}

func main() {
	bestRunner := meters(10)
	fmt.Printf("%0.3fm/s equals %0.3f miles/s\n", bestRunner, bestRunner.Tomiles())
	toSeoul := kilometers(60)
	fmt.Printf("%0.3fkm equals %0.3f miles\n", toSeoul, toSeoul.Tomiles())
}
