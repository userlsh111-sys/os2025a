package main

import (
	"fmt"
	"log" //현재 시각 출력
	"studyCapsulEmb/pkg/calendar"
)

func main() {
	today := calendar.Date{}
	err := today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetMonth(16)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today.RtYear(), "년 ", today.RtMonth(), "월 ", today.RtDay(), "일")
}
