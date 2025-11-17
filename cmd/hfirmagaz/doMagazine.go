package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
	//타입을 대신 선언해준 magazine.go
)

func main() {
	var j1 magazine.Employee
	var p1 magazine.Address
	j1.Name = "Jamyeong Lee"
	j1.Salary = 200000.0
	p1.Street = "Itaewon"
	p1.City = "Yongsan, Seoul"
	fmt.Println(j1.Name)
	fmt.Println(j1.Salary)
	fmt.Println(p1.Street)
	fmt.Println(p1.City)
}
