package main

import "fmt"

type Car struct {
	brand string
	year  int16
}

func structsUse() {
	myCar := Car{brand: "BMW", year: 2003}
	fmt.Println(myCar)

	// Other way to instnace
	var otherCar Car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
