package main

import "fmt"

func chain(left, right chan int) {
	left <- 1 + <-right
}
func main() {
	fmt.Println("Daisy Chain")

	leftMost := make(chan int)
	right := leftMost
	left := leftMost
	for i := 0; i < 100; i++ {
		right = make(chan int)
		go chain(left, right)
		left = right
	}

	go func(ch chan int) {
		ch <- 1
	}(right)
	fmt.Println(<-leftMost)
}
