package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findingArea(width float64, height float64) float64 { //return값의 type: float64
	area := width * height
	return area
}
func findShowArea() {
	fmt.Print("The next wid: ")
	wid := bufio.NewReader(os.Stdin)
	rwid, err := wid.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	realwid := strings.TrimSpace(rwid)               //여까진 str로 잘 받음.
	rrealwid, err := strconv.ParseFloat(realwid, 64) //이제 float64로.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("The next hei: ")
	hei := bufio.NewReader(os.Stdin)
	rhei, err := hei.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	realhei := strings.TrimSpace(rhei)               //여까진 str로 잘 받음.
	rrealhei, err := strconv.ParseFloat(realhei, 64) //이제 float64로.
	if err != nil {
		log.Fatal(err)
	}

	area := findingArea(rrealwid, rrealhei)
	fmt.Printf("The next area is %.2f.", area)
}

func main() {
	fmt.Printf("Just print(2.525, 5.555): %.2f\n", findingArea(2.525, 5.555))

	findShowArea()
}
