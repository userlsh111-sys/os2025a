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
	input, err := read.ReadString('\n') // ignore error(빈 식별자)
	log.Fatal(err)                      //error 보고, 종료
	fmt.Println(input)
}
