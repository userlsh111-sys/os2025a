package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

func main() {
	start := time.Now()
	go say("고 루틴")
	say("메인")
	//go 전: 둘이 함께.
	//go 후: 하나 끝나고, 다른 거 하나.

	fmt.Println("실행 시간:", time.Since(start))
}
