package main

import (
	"fmt"
	"strings"
)

func isPalindrom(text string) {
	strings.ToLower(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("is palindrom")
	} else {
		fmt.Println("is not palindrom")
	}
}
func arraySliceUse() {
	// Array
	var array [5]int
	array[0] = 1
	array[1] = 2

	// Cap get may capacity of arrays or eslices
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metethos
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice))

	// Append list
	newSlice := []int{8, 9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))

	for i, value := range slice {
		fmt.Println(i, value)
	}
	for i := range slice {
		fmt.Println(i)
	}
	for _, value := range slice {
		fmt.Println(value)
	}

	isPalindrom("casa")
	isPalindrom("amor a roma")
	isPalindrom("Ama")
	isPalindrom("perro")
}
