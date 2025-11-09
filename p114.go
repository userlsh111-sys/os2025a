package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("Enter the score: ")
	score, err := getFloat() //getFloat() 호출
	if err != nil {
		fmt.Print("mistaken bufio, err")
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "pass, good job."
	} else {
		status = "fail, sorry."
	}
	fmt.Print("Your score", score, ",", status)
}
