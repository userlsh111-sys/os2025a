package main

import "fmt"

func main() {
	reaptStr("영구꾸 ハムスター", "화이팅", 4)
}

func reaptStr(fir string, sec string, times int) {
	for i := 0; i < times; i++ {
		if i%2 == 0 {
			fmt.Printf("%s\n", fir)
		} else {
			fmt.Printf("%s\n", sec)
		}
	}
}
