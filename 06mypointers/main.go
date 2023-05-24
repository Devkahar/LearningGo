package main

import "fmt"

func main() {
	fmt.Println("We will be using pointers")

	// var ptr *int
	// fmt.Println("Value of pointer ", ptr)

	actualValue := 23

	var ptr = &actualValue
	fmt.Println("This is actual Value  ", ptr)
	fmt.Println("This is actual Value  ", *ptr)
	*ptr = *ptr * 2
	fmt.Println("This is actual Value after operation:  ", actualValue)

}
