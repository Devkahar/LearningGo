package main

import "fmt"

func main() {
	fmt.Println("Learning function in go")
	greeting()
	sum := adder(2, 4)
	fmt.Println("Sum 2 and 4 ", sum)
	total, message := proSum(1, 2, 3, 4, 5)

	fmt.Println("Total from pro is: ", total)
	fmt.Println("Message from pro is: ", message)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proSum(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hey i am form pro"
}

func greeting() {
	fmt.Println("Namaste from GO")
}
