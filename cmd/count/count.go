package main

import (
	"fmt"
	"log"
	"studyMap/pkg/datafile"
)

func main() {
	lines, err := datafile.GetStrings("candidates.txt")
	//외부 패키지의 것을 가져올 땐 첫 글자가 대문자여야..
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
