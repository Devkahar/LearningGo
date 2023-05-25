package main

import "fmt"

func main() {
	fmt.Println("Learing Defer in Go")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
