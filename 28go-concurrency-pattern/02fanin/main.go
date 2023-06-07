package main

import (
	"fmt"
	"time"
)

func main() {
	// Fanin also know as multiplexing.
	// Speaker-1
	// 				-> Speaker("Listning to who ever speak")
	// Speaker-2

	c := fanIn(getGreeting("Dev"), getGreeting("Ana"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

}

func fanIn(bob, ana <-chan string) <-chan string {
	ch := make(chan string)

	// go func() {
	// 	for {
	// 		ch <- <-bob
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		// Lets say ana take long time to reply
	// 		time.Sleep(2100 * time.Millisecond)
	// 		ch <- <-ana
	// 	}
	// }()

	// --------------------------------
	// using select statement for same.

	go func() {
		for {
			select {
			case s := <-bob:
				ch <- s
			case s := <-ana:
				time.Sleep(1 * time.Millisecond)
				ch <- s
			}
		}
	}()

	return ch
}

func getGreeting(msg string) <-chan string {

	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("Hey %s, Good Morning", msg)
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}
