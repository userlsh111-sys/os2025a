package main

import (
	"fmt"
	"log"
	"studyPkg/pkg/keyboard"
)

func main() {
	fmt.Print("Enter the score: ")
	score, err := keyboard.GetFloat()
	//keyboard 패키지의 getFloat() 호출
	if err != nil {
		fmt.Print("mistaken bufio, err")
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "pass, good job"
	} else {
		status = "fail, sorry"
	}
	fmt.Printf("Your score=%.1f, %s.", score, status)
}
