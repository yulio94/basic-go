package main

import "fmt"

func deferUse() {

	defer fmt.Println(",World")
	fmt.Println("Hello")
}
