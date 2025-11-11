package main

import "fmt"

func main() {
	reaptStr("이재명", "탄핵", 3)
}

func reaptStr(fir string, sec string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%s %s\n", fir, sec)
	}
}
