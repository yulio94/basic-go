package main

import (
	"fmt"
)

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

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	// Enteros positivos
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	examplesFmt()
	mainTestFunctions()
	showBucle()
	deferUse()
	continueBreakUse()
	arraySliceUse()
	mapsUse()
	structsUse()
	pointersUse()
	interfacesUse()
	goRoutinesUse()
}
