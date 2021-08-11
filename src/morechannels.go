package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func moreChannelsUse() {
	c := make(chan string, 2)
	c <- "Message one"
	c <- "Message two"

	fmt.Println(len(c), cap(c))

	// Range and close
	// Close channels and avoid more insertions
	close(c)

	// range iterate over message in a channel
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email := make(chan string)
	emailTwo := make(chan string)

	go message("j@email.com", email)
	go message("c@emai.com", emailTwo)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("email one", m1)
		case m2 := <-emailTwo:
			fmt.Println("email two", m2)
		}
	}

}
