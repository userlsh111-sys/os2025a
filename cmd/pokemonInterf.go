package main

import "fmt"

// Pokemon 인터페이스 정의
type Pokemon interface {
	Name() string
	Type() string
	Attack() int
	Defense() int
}

type Lizard struct {
	hp int
}

func (c Lizard) Name() string {
	return "리자드"
}
func (c Lizard) Type() string {
	return "불꽃"
}
func (c Lizard) Attack() int {
	return 52
}
func (c Lizard) Defense() int {
	return 43
}

type Kkbk struct {
	hp int
}

func (s Kkbk) Name() string {
	return "꼬부기"
}
func (s Kkbk) Type() string {
	return "물"
}
func (s Kkbk) Attack() int {
	return 48
}
func (s Kkbk) Defense() int {
	return 65
}

type Strange struct {
	hp int
}

func (b Strange) Name() string {
	return "이상해씨"
}
func (b Strange) Type() string {
	return "풀"
}
func (b Strange) Attack() int {
	return 49
}
func (b Strange) Defense() int {
	return 49
}

// Pokemon을 파라미터로 받는 함수
func printPokemonInfo(p Pokemon) {
	fmt.Printf("%s (%s)\n", p.Name(), p.Type())
	fmt.Printf("공격: %d, 방어: %d\n\n",
		p.Attack(), p.Defense())
}

func main() {
	lizard := Lizard{hp: 39}
	kkbk := Kkbk{hp: 44}
	bulbasaur := Strange{hp: 45}

	// 모두 Pokemon 인터페이스로 사용 가능
	printPokemonInfo(lizard)
	printPokemonInfo(kkbk)
	printPokemonInfo(bulbasaur)
}
