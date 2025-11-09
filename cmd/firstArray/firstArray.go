package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayBool := [3]bool{true, false, true} //리터럴
	var arrayInt [3]int
	fmt.Println(arrayBool[1])
	fmt.Println(reflect.TypeOf(arrayBool))

	fmt.Println("아래: arrayInt 1번 인덱스 비교")
	fmt.Println(arrayInt[1])
	arrayInt[1]++
	arrayInt[1] = arrayInt[1] + 1
	fmt.Println(arrayInt[1])

	fmt.Printf("%#v\n", arrayInt)
	fmt.Println("혹은 for로 출력")
	for i := 0; i < 3; i++ {
		fmt.Printf("인덱스 %d 요소: %d\n", i, arrayInt[i])
	}
}
