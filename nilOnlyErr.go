package main

import (
	"bufio"
	"fmt"
	"log" //+error 처리
	"os"
)

func main() {
	fmt.Print("성적 입력: ")
	read := bufio.NewReader(os.Stdin)
	input, err := read.ReadString('\n') //ignore error(빈 식별자)
	if err != nil {
		log.Fatal(err)
	} //조건문으로 error 보고 및 종료 제한
	fmt.Println(input)
}
