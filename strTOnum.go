package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("성적: ")
	read := bufio.NewReader(os.Stdin)
	input, err := read.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	//status 변수 선언
	//조건문 안의 :(선언)은 의미 없음.

	fmt.Println("결과: ", status)
}
