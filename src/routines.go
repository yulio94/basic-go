package main

import (
	"fmt"
	"sync"
)

func say(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(message)
}

func goRoutinesUse() {
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)
	go say("World", &wg)
	wg.Wait()
	go func(text string) {
		fmt.Println(text)
	}("Bye")
}
