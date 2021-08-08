package main

import "fmt"

func examplesFmt() {
	worldMessage := "World"
	helloMessage := "Hello"

	//Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	name := "Platzi"
	var courses uint16 = 500
	fmt.Printf("%s have more than %d\n", name, courses)
	// When we don't know the variable type use %v
	fmt.Printf("%v have more than %v\n", name, courses)

	// Sprintf
	message := fmt.Sprintf("%s have more than %d", name, courses)
	fmt.Println(message)

	//Data type
	fmt.Printf("Hello message: %T\n", helloMessage)
	fmt.Printf("Courses: %T\n", courses)
}
