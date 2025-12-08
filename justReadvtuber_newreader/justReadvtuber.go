package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	stellive := []string{"yuni", "kangji", "who"}
	aesther := []string{"elly", "karin"}
	veuyai := []string{"mei"}
	preojecti := []string{"honey churros"}
	German := []string{"shylily"}
	ekis := []string{"miorri"}

	outputs := []string{
		"Be glaaaabal.",
		"I approve.",
		"Wrong input.",
		"Fxxxxxxxxxxxxxxxxxxxk you.",
		"sachoooo",
	}

	fmt.Println("Favorite Vtuber(first name):")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') //낚싯줄(\n 자리잡음)
	if err != nil {
		log.Fatal(err)
	} //어쩌다 nil 복습..
	rlinput := strings.ToLower(input)
	realinput := strings.TrimSpace(rlinput) //칼(\n 삭제)
	fmt.Println("Your answer:", realinput)

	var output string

	if realinput == stellive[0] || realinput == veuyai[0] {
		output = outputs[1] //전역과 지역 변수처럼 블록 밖과 안도 다름.
	} else if realinput == stellive[1] {
		output = outputs[4]
	} else if realinput == stellive[2] ||
		realinput == aesther[0] || realinput == aesther[1] {
		output = outputs[0]
	} else if realinput == preojecti[0] {
		fmt.Println("Sponsered by shy..")
	} else if realinput == German[0] {
		output = outputs[1]
	} else if realinput == ekis[0] {
		output = outputs[3]
	} else {
		output = outputs[2]
	}

	fmt.Println("My answer:", output)
}
