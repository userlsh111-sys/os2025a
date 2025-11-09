package main

import (
	"fmt"
	"log"
	"studyPkg/pkg/keyboard"
)

func main() {
	fmt.Print("Enter the F temp.: ")
	fahrenheit, err := keyboard.GetFloat()
	//keyboard 패키지의 getFloat() 호출
	if err != nil {
		fmt.Print("mistaken bufio, err")
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("0.2f degrees Celsius\n", celsius)
}
