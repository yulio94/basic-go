package main

import (
	"fmt"
)

func mapsUse() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Sharon"] = 23

	fmt.Println(m)

	// Iterate throw map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Find value
	value := m["Jose"]
	fmt.Println(value)

	wrongValue, ok := m["Josep"]
	fmt.Println(wrongValue, ok)
}
