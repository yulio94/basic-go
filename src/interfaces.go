package main

import "fmt"

type twoDimensionFigure interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (sq square) area() float64 {
	return sq.base * sq.base
}

func (rec rectangle) area() float64 {
	return rec.base * rec.height
}

func calculate(f twoDimensionFigure) {
	fmt.Println("Area: ", f.area())
}

func interfacesUse() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 5}

	calculate(mySquare)
	calculate(myRectangle)

	// List of interfaces
	myInterfaceList := []interface{}{"Hello", 1, true, ", world"}
	fmt.Println(myInterfaceList...)
}
