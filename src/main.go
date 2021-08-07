package main

import "fmt"

func main() {
	// Define constants
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("pi", pi)
	fmt.Println("pi 2", pi2)

	//Define variables
	base := 12
	var height int = 14
	var area int
	fmt.Println(base, height, area)

	// Zero value
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("int", a, "float64", b, "string", c, "bool", d)

	// Area of a square
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Area of a square", squareArea)
}
