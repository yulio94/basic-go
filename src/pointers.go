package main

import "fmt"

type player struct {
	first_name string
	last_name  string
	age        uint8
	gender     string
}

// Without pointer to struct don't allow modified struct instance values.
func (player player) getFullName() string {
	return player.last_name + player.last_name
}

// With pointer to struct allow modified struct instance values.
func (player *player) updateFullName(first_name string, last_name string) string {
	player.first_name = first_name
	player.last_name = last_name
	return player.last_name + player.last_name
}

func pointersUse() {
	var a uint8 = 50
	var b = &a

	// * get memory direction of a pointer
	fmt.Println(&a)
	// * get memory direction of a pointer
	fmt.Println(b)
	// * get value of a pointer
	fmt.Println(*b)

	var creature string = "shark"
	var pointer *string = &creature
	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	messi := player{first_name: "Lionel", last_name: "Messi", age: 33, gender: "Male"}
	fmt.Println("GOAT: ", messi.getFullName())
	fmt.Println("GOAT: ", messi)

	newName := messi.updateFullName("Lionel Andres", "Messi Puchitiqui")
	fmt.Println("GOAT: ", newName)
	fmt.Println("GOAT: ", messi)
}
