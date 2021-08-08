package main

import "fmt"

func continueBreakUse() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Printf("I have %d number \n", i)
			continue
		} else if i == 8 {
			fmt.Printf("Finish on: %d \n", i)
			break
		} else {
			fmt.Println(i)
		}

	}
}
