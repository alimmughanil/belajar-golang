package main

import "fmt"

func main() {
	getVariable()
	getMathOperator()
	getCondition()
	getLogic(true, true)
	getLooping(5)
}

func getVariable() {
	var ageHuman int = 10
	ageHuman = 20
	phi := 3.14
	fmt.Println(ageHuman, phi)
}

func getMathOperator() {
	s := 4
	L := s * s

	fmt.Println(L)
}

func getCondition() {
	hasil := 5 > 4
	fmt.Println(hasil)
}

func getLogic(a bool, b bool) {
	hasil := a && b
	fmt.Println(hasil)
}

func getLooping(maxNumber int) {
	for i := 1; i <= maxNumber; i++ {
		result := fmt.Sprintf("%d%s", i, ". Hello World")
		fmt.Println(result)
	}
}
