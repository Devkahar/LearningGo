package main

import "fmt"

func main() {
	fmt.Println("Learning if/else in GO")

	if 7%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
	if num := 3; num > 10 {
		fmt.Println("Num is > 10")
	} else {
		fmt.Println("Num is < 10")
	}
}
