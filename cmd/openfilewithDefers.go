package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("testHime.txt")
	if err != nil {
		fmt.Println("오픈 실패:", err)
		return
	}

	defer file.Close() //닫기 예약

	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("읽는 중...")
	//순서가 stack형이다.
}
