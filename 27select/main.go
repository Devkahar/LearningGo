package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println("Func Called")
	x, y := 0, 1
	for {
		fmt.Println("In Func Loop")
		select {
		case c <- x:
			fmt.Println("Push value in ch ", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			fmt.Println("In Loop", i)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// ----------- Select Statement ---------
// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case.
// It chooses one at random if multiple are ready.
//
// ----------- Default Case in Select -----------
// The default case in a select is run if no other case is ready.
// Use a default case to try a send or receive without blocking:
// select {
// case i := <-c:
//
//	// use i
//
// default:
//
//	    // receiving from c would block
//	}
