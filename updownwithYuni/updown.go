package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	neededSeed := time.Now().Unix()
	rand.Seed(neededSeed)
	answer := rand.Intn(100) + 1

	justreader := bufio.NewReader(os.Stdin)

	whether := false

	for howmany := 1; howmany <= 10; howmany++ {

		fmt.Println("?: 이루뿌터 빼까찌 중에 수짜루르 마춰빠~")
		guess, err := justreader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		convedGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if convedGuess >= 1 && convedGuess <= 100 {
			if howmany == 10 {
				break
			} else if convedGuess < answer {
				fmt.Println("안 작다고.... 비질게.(남은 기회:", 10-howmany, ")")
				//printf로 변경 요망
			} else if convedGuess > answer {
				fmt.Println("'62'보다 크네!(남은 기회:", 10-howmany, ")")
			} else {
				whether = true
				fmt.Println("(대충 디프츄츄 짤)")
				break
			}
		} else if convedGuess <= 1 {
			fmt.Println("자신가믈 가져! 유니쳐럼!")
		} else if convedGuess >= 100 {
			fmt.Println("빼꾸보다 작따고. 존나 바보야 www")
		} else {
			fmt.Println("(1~100 사이의 정수를 입력하세요.)")
		}
	}

	if !whether {
		fmt.Println("빠꾸대가리!")
	}
}
