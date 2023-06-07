package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Generator Pattern")
	// In Generator pattern we return channel from function

	ch := getGreeting("Dev")

	for i := 0; i < 5; i++ {
		fmt.Printf("Recive Greeting %d -> %s\n", i+1, <-ch)
	}
}

// getGreeting is a generator function that return a channel.
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
