package main

import "fmt"

func saySomething(text string, c chan<- string) {
	c <- text
}
func channelsUse() {
	c := make(chan string, 1)

	fmt.Println("Hello")
	go saySomething("Bye", c)
	fmt.Println(<-c)
}
