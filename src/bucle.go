package main

import "fmt"

func showBucle() {

	// Conditional for
	for i := 0; i <= 10; i++ {
		fmt.Printf("Cicle: %d \n", i)
	}

	// Fow while
	counter := 0
	for counter < 10 {
		fmt.Printf("Cicle while: %d \n", counter)
		counter++
	}

	// For forever
	// counterForever := 0
	// for {
	// 	fmt.Printf("Cicle while forever: %d \n", counterForever)
	// 	counter++
	// }

}
