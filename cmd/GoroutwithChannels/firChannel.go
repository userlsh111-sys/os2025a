package main

import (
	"fmt"
	"time"
)

func Virtuber(channel chan string) {
	channel <- "Yuni"
	channel <- "Karin"
	channel <- "Mei"
}

func Youtuber(channel chan string) {
	channel <- "Mrbeast"
	channel <- "Benio"
	channel <- "Ratevfx"
	channel <- "Gooseworx"
}

func main() {
	time.Sleep(1000 * time.Millisecond)
	start := time.Now()

	aya := make(chan string)
	beni := make(chan string)
	go Virtuber(aya)
	go Youtuber(beni)
	fmt.Println(<-beni)
	fmt.Println(<-aya)
	fmt.Println(<-beni)
	fmt.Println(<-aya)
	fmt.Println(<-beni)
	fmt.Println(<-aya)
	fmt.Println(<-beni)

	fmt.Println("실행 시간:", time.Since(start))
}

//교수님은 kbs, mbc에 비유함. 채널을 돌린다고 다른 채널이 바뀌진 않지.
