package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func rateofDesign(design float64, full float64) float64 {
	rate := design / full
	return rate * 100
}

func main() {
	fmt.Println("Full charge:")
	designer := bufio.NewReader(os.Stdin)
	rinput1, err := designer.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	realinput1 := strings.TrimSpace(rinput1)
	fmt.Println("Your 'full':", realinput1)

	rrealinput1, err := strconv.ParseFloat(realinput1, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Design charge:")
	fullP := bufio.NewReader(os.Stdin)
	rinput2, err := fullP.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	realinput2 := strings.TrimSpace(rinput2)
	fmt.Println("Your 'design':", realinput2)

	rrealinput2, err := strconv.ParseFloat(realinput2, 64)
	if err != nil {
		log.Fatal(err)
	}

	percents := rateofDesign(rrealinput1, rrealinput2)
	fmt.Println("Rate:", percents)
}
