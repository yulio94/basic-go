package main

import "fmt"

func returnOneValue(a uint) uint {
	return a
}

func returnTwoValues(a uint, b uint) (result, resultTwo uint) {
	return result, resultTwo
}

func mainTestFunctions() {
	var number uint = 10
	var numberTwo uint = 20
	result := returnOneValue(number)
	resultTwo, resultThree := returnTwoValues(number, numberTwo)
	resultWithIgnoreOne, _ := returnTwoValues(number, numberTwo)

	fmt.Printf("Los resultados son: %d %d %d %d \n", result, resultTwo, resultThree, resultWithIgnoreOne)
}
